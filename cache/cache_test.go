package cache

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultCache(t *testing.T) {
	t.Parallel()

	// given
	cache := NewLocalCache(context.Background(), nil)

	type Info struct {
		Id   uint64 `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	err := cache.Set("cache:user:id", Info{
		Id:   100,
		Name: "nash",
	})
	assert.Equal(t, err, nil)

	res := &Info{}
	err = cache.Get("cache:user:id", res)
	assert.Equal(t, err, nil)

	assert.Equal(t, res.Id, uint64(100))
	assert.Equal(t, res.Name, "nash")

}

func TestConfigCache(t *testing.T) {
	t.Parallel()

	// 1G
	cache := NewLocalCache(context.Background(), NewMemoryConfig(SetMaxMemorySize(1024)))

	type Info struct {
		Id   uint64 `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	err := cache.Set("cache:user:id", Info{
		Id:   100,
		Name: "nash",
	})
	assert.Equal(t, err, nil)

	res := &Info{}

	err = cache.Get("cache:user", res)
	assert.Equal(t, err, Nil)

	err = cache.Get("cache:user:id", res)
	assert.Equal(t, err, nil)

	assert.Equal(t, res.Id, uint64(100))
	assert.Equal(t, res.Name, "nash")

}
