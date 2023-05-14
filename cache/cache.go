package cache

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"time"

	"github.com/allegro/bigcache/v3"
	"go.uber.org/zap"
)

type MemoryFunc func(t *MemoryConfig)

// MemoryConfig 缓存配置
type MemoryConfig = bigcache.Config

func NewMemoryConfig(opts ...MemoryFunc) *MemoryConfig {

	conf := &MemoryConfig{
		Shards:             1024,             // Number of cache shards, value must be a power of two
		LifeWindow:         10 * time.Minute, // life time
		CleanWindow:        5 * time.Minute,  // clean time
		MaxEntriesInWindow: 1000 * 10 * 60,   // qps * lifeWindow = 生命周期内最多的entry数量
		MaxEntrySize:       500,              // unit : Byte  单个entry的大小
		StatsEnabled:       false,
		Verbose:            false,
		Hasher:             nil,
		HardMaxCacheSize:   50, // unit: MB
		Logger:             nil,
	}

	for _, opt := range opts {
		opt(conf)
	}

	return conf
}

// SetMaxMemorySize 设置最大物理内存(uint: MB)
func SetMaxMemorySize(hardCacheSize int) MemoryFunc {
	return func(t *MemoryConfig) {
		t.HardMaxCacheSize = hardCacheSize
	}
}

// SetLifeWindow 设置对象生命周期
func SetLifeWindow(life time.Duration) MemoryFunc {
	return func(t *MemoryConfig) {
		t.LifeWindow = life
	}
}

// SetCleanWindow 设置清理时间间隔
func SetCleanWindow(cleanInterval time.Duration) MemoryFunc {
	return func(t *MemoryConfig) {
		t.CleanWindow = cleanInterval
	}
}

// SetMaxEntrySize 基于qps计算生命周期内最大成员数量
func SetMaxEntrySize(qps int) MemoryFunc {
	return func(t *MemoryConfig) {
		t.MaxEntriesInWindow = qps * int(t.LifeWindow.Seconds())
	}
}

// LocalCache 本地缓存
type LocalCache struct {
	cache *bigcache.BigCache
}

func NewLocalCache(ctx context.Context, memConfig *MemoryConfig) *LocalCache {

	defaultConfig := bigcache.DefaultConfig(10 * time.Minute)
	if memConfig != nil {
		defaultConfig = *memConfig
	}

	cache, err := bigcache.New(ctx, defaultConfig)
	if err != nil {
		log.Printf("NewLocalCache:%+v", zap.Any("NewLocalCache", err))
		return nil
	}
	return &LocalCache{
		cache: cache,
	}
}
func (t *LocalCache) Get(key string, dst interface{}) error {
	res, err := t.cache.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return Nil
	}

	if err != nil {
		log.Print(zap.Any("LocalCache Get err", err))
		return err
	}

	if reflect.TypeOf(dst).Kind() != reflect.Ptr {
		log.Print(zap.Any("dst must pointer type", err))
		return ErrDataTypeInvalid
	}
	return json.Unmarshal(res, dst)
}

func (t *LocalCache) Set(key string, dst interface{}) error {
	data, err := json.Marshal(dst)
	if err != nil {
		return err
	}
	return t.cache.Set(key, data)
}
