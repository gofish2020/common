package gredis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewClient() *Client {
	return &Client{
		wrapRedisClient{
			rdb: redis.NewClient(&redis.Options{
				Addr:     "localhost:36379",
				Password: "G62m50oigInC30sf", // no password set
				DB:       0,                  // use default DB
			}),
		},
	}
}

type Client struct {
	wrapRedisClient
}

type wrapRedisClient struct {
	rdb *redis.Client
}

func (t *wrapRedisClient) Set(ctx context.Context, key string, value interface{}, expire time.Duration) (string, error) {
	return t.rdb.Set(ctx, key, value, expire).Result()

}

func (t *wrapRedisClient) Get(ctx context.Context, key string) (string, error) {
	return t.rdb.Get(ctx, key).Result()
}

func (t *wrapRedisClient) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return t.rdb.HSet(ctx, key, values...).Result()
}

func (t *wrapRedisClient) Publish(ctx context.Context, channel string, msg interface{}) (int64, error) {
	return t.rdb.Publish(ctx, channel, msg).Result()
}

func (t *wrapRedisClient) Subscripe(ctx context.Context, channle ...string) *redis.PubSub {

	return t.rdb.Subscribe(ctx, channle...)
}

func (t *wrapRedisClient) MSet(ctx context.Context, values ...interface{}) (string, error) {
	return t.rdb.MSet(ctx, values...).Result()
}

func (t *wrapRedisClient) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return t.rdb.MGet(ctx, keys...).Result()
}
