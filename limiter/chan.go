package limiter

import "sync/atomic"

/*
	基于channel实现的限流器
*/

const (
	DefaultLimit = 10
)

type ConcurrencyLimiter struct {
	tickets chan int
	limit   int
	running int32
}

func NewConcurrencyLimiter(limit int) *ConcurrencyLimiter {
	if limit < 0 {
		limit = DefaultLimit
	}

	tickets := make(chan int, limit)
	for i := 0; i < limit; i++ {
		tickets <- i
	}

	return &ConcurrencyLimiter{
		limit:   limit,
		tickets: tickets,
	}
}

func (t *ConcurrencyLimiter) Execute(task func()) int {
	ticket := <-t.tickets
	atomic.AddInt32(&t.running, 1)
	go func() {
		defer func() {
			t.tickets <- ticket
			atomic.AddInt32(&t.running, -1)
		}()
		task()
	}()
	return ticket
}

func (t *ConcurrencyLimiter) Wait() {
	for i := 0; i < t.limit; i++ {
		<-t.tickets
	}
}

func (t *ConcurrencyLimiter) RunningNums() int32 {
	return t.running
}
