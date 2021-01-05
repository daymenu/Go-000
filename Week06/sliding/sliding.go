package sliding

import (
	"sync"
	"time"
)

const defaultWindowTime = 1 * time.Second
const defaultBucketTime = 200 * time.Millisecond
const defaultMaxBucket = int(200 * defaultWindowTime / defaultBucketTime)
const defaultMixBucket = int(5 * defaultWindowTime / defaultBucketTime)

type slidingOptions struct {
	windowTime time.Duration
	bucketTime time.Duration
}

var defaultSlidingOption = slidingOptions{
	windowTime: defaultWindowTime,
	bucketTime: defaultBucketTime,
}

// CounterOption define a sliding option interface
type CounterOption interface {
	apply(*slidingOptions)
}

type funcCounterOption struct {
	f func(*slidingOptions)
}

func (fdo *funcCounterOption) apply(do *slidingOptions) {
	fdo.f(do)
}

func newFuncCounterOption(f func(*slidingOptions)) *funcCounterOption {
	return &funcCounterOption{
		f: f,
	}
}

// WindowTime set a window time
func WindowTime(wt time.Duration) CounterOption {
	return newFuncCounterOption(func(o *slidingOptions) {
		o.windowTime = wt
	})
}

// BucketTime  set a bucket time
func BucketTime(bt time.Duration) CounterOption {
	return newFuncCounterOption(func(o *slidingOptions) {
		o.bucketTime = bt
	})
}

// Counter counter
type Counter struct {
	sync.RWMutex
	opts  slidingOptions
	world map[int64]*bucket
}

type bucket struct {
	success int64
	time    int64
}

// NewCounter new a sliding window counter
func NewCounter(opt ...CounterOption) *Counter {
	opts := defaultSlidingOption

	for _, o := range opt {
		o.apply(&opts)
	}

	c := &Counter{
		opts:  opts,
		world: make(map[int64]*bucket),
	}

	return c
}

// Inc increment count success
func (c *Counter) Inc(i int64) {
	c.Lock()
	l := len(c.world)
	if l > defaultMaxBucket {
		t := time.Now().UnixNano() - int64(defaultMixBucket*int(c.opts.bucketTime.Nanoseconds()))
		clearKey := c.getBucketKey(t)
		for k := range c.world {
			if k < clearKey {
				delete(c.world, k)
			}
		}
	}
	t := c.getBucketKey(time.Now().UnixNano())
	if b, ok := c.world[t]; ok {
		b.success += i
		c.Unlock()
		return
	}
	b := bucket{success: i, time: t}
	c.world[t] = &b
	c.Unlock()
}

func (c *Counter) getBucketKey(t int64) int64 {
	return t / c.opts.bucketTime.Nanoseconds()
}

func (c *Counter) currentBucket(t int64) *bucket {
	key := c.getBucketKey(t)
	c.RLock()
	if b, ok := c.world[key]; ok {
		c.RUnlock()
		return b
	}
	c.RUnlock()
	b := bucket{success: 0, time: key}
	return &b
}

func (c *Counter) windowBuckets(t int64) []*bucket {
	var buckets []*bucket
	rWin := time.Now().UnixNano()
	lWin := rWin - int64(c.opts.windowTime)
	for lWin <= rWin {
		buckets = append(buckets, c.currentBucket(lWin))
		lWin += c.opts.bucketTime.Nanoseconds()
	}
	return buckets
}

// Sum 窗口内总数
func (c *Counter) Sum(t time.Time) (sum int64) {
	buckets := c.windowBuckets(t.UnixNano())
	for _, b := range buckets {
		sum += b.success
	}
	return sum
}

// Avg 每个bucket内的平均数
func (c *Counter) Avg(t time.Time) int64 {
	return c.Sum(t) / (int64(c.opts.windowTime) / int64(c.opts.bucketTime))
}
