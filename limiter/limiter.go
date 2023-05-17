package limiter

import (
	"time"

	"github.com/juju/ratelimit"
)

/*
	基于第三方包 github.com/juju/ratelimit@latest
*/

type RateLimiter struct {
	bucket *ratelimit.Bucket
}

// fillInterval 令牌填充间隔【每次填充1个令牌】 capacity 令牌最大容量
func NewRateLimiter(fillInterval time.Duration, capacity int64) *RateLimiter {
	return &RateLimiter{
		bucket: ratelimit.NewBucket(fillInterval, capacity),
	}
}

// fillInterval 令牌填充间隔【每次填充quantum个令牌】 capacity 令牌最大容量
func NewRateLimiterWithQuantum(fillInterval time.Duration, capacity, quantum int64) *RateLimiter {
	return &RateLimiter{
		bucket: ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum),
	}
}

// rate 每秒填充rate个令牌【类似于qps】 capacity 令牌最大容量
func NewBucketWithRate(rate float64, capacity int64) *RateLimiter {
	return &RateLimiter{
		bucket: ratelimit.NewBucketWithRate(rate, capacity),
	}
}

// Wait 阻塞获取count个令牌
func (t *RateLimiter) Wait(count int64) {
	t.bucket.Wait(count)
}

// WaitMaxDuration 最长阻塞等待maxWait，如果没有获取到返回false
func (t *RateLimiter) WaitMaxDuration(count int64, maxWait time.Duration) bool {
	return t.bucket.WaitMaxDuration(count, maxWait)
}

// Take 非阻塞获取count个令牌，返回值表示：如果获取count个令牌，需要等待的时长【=0】表示获取到了count个令牌， 【>0】只获取了部分令牌，剩下的令牌需要等待time.Duration
func (t *RateLimiter) Take(count int64) time.Duration {
	return t.bucket.Take(count)
}

// TakeMaxDuration 获取count个令牌需要的时长> maxWait，相当于什么也不执行bool=false; 否则获取到【0-count】个令牌，以及等待的时长，类似于Take
func (t *RateLimiter) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool) {
	return t.bucket.TakeMaxDuration(count, maxWait)
}

// TakeAvailable 非阻塞获取count个令牌，返回值为实际得到的令牌数量
func (t *RateLimiter) TakeAvailable(count int64) int64 {
	return t.bucket.TakeAvailable(count)
}

// Available 当前可用的令牌数量【只用于调试】不能作为判断可用令牌的依据
func (t *RateLimiter) Available() int64 {
	return t.bucket.Available()
}

// Rate 令牌的填充速率 , unit : 【个/秒】
func (t *RateLimiter) Rate() float64 {
	return t.bucket.Rate()
}
