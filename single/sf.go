package single

import (
	"context"
	"log"

	"github.com/gofish2020/common/cache"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

type InterFunc func(ctx context.Context) (interface{}, error)

type AwesomeCache struct {
	sf         singleflight.Group
	localCache *cache.LocalCache
	ctx        context.Context
}

func NewAwesomeCache(ctx context.Context, memoryConfig *cache.MemoryConfig) *AwesomeCache {
	return &AwesomeCache{
		sf:         singleflight.Group{},
		localCache: cache.NewLocalCache(ctx, memoryConfig),
		ctx:        ctx,
	}
}

func (t *AwesomeCache) Get(key string, f InterFunc) (interface{}, error) {

	var result = new(interface{})

	if t.localCache.Get(key, result) == cache.ErrEntryNotFound { // 未命中缓存
		data, err, _ := t.sf.Do(key, func() (interface{}, error) {

			log.Printf("NewAwesomeCache:%+v\n", zap.String("NewAwesomeCache->Get", "get from db"))
			dbData, err := f(t.ctx)
			if err != nil {
				return nil, err
			}
			t.localCache.Set(key, dbData) //设置缓存
			return dbData, nil
		})

		if err != nil {
			return nil, err
		}
		return data, nil // 从singlefight中获取的结果
	}
	return *result, nil
}
