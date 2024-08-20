package medium

import (
	"fmt"
)

type CacheKey int

type LFUCache struct {
	Cache   map[CacheKey]int
	Counter map[CacheKey]int64
	Size    int
}

func NewLFUCache(size int) *LFUCache {
	return &LFUCache{
		Cache:   make(map[CacheKey]int),
		Counter: make(map[CacheKey]int64),
		Size:    size,
	}
}

func (c *LFUCache) UsedSpace() int {
	return len(c.Cache)
}

func (c *LFUCache) IsFull() bool {
	return c.UsedSpace() >= c.Size
}

// Get min frequency key in cache
func (c *LFUCache) GetMinFreqKey() CacheKey {
	minFreq := int64(0)
	smallestKey := CacheKey(0)
	for key, freq := range c.Counter {
		if minFreq == 0 || freq < minFreq {
			minFreq = freq
			smallestKey = key
		} else if minFreq == freq && key < smallestKey {
			smallestKey = key
		}
	}

	return smallestKey
}

func (c *LFUCache) Get(key CacheKey) int {
	val, ok := c.Cache[key]
	if ok {
		return val
	}
	return -1
}

func (c *LFUCache) Update(key CacheKey, value int) {
	_, ok := c.Cache[key]
	if ok {
		// * in case key is already exist, just update and increase counter
		c.Cache[key] = value
		c.Counter[key] += 1
		return
	}

	if c.IsFull() {
		// * in case cache is full, invalidate min freq cache before add new
		keyToDelete := c.GetMinFreqKey()
		if keyToDelete != 0 {
			c.Delete(keyToDelete)
		}
	}

	c.Cache[key] = value
	c.Counter[key] = 1
}

// clear cache also clear counter of cached key
func (c *LFUCache) Delete(key CacheKey) {
	delete(c.Cache, key)
	delete(c.Counter, key)
}

type OpType int

const (
	Get    OpType = 1
	Update OpType = 2
)

type Operation struct {
	Type  OpType
	Key   int
	Value int
}

func OperateCache(c *LFUCache, op Operation) {
	switch op.Type {
	case Get:
		fallthrough
	case Update:
		fallthrough
	default:
		panic("Wrong operation")
	}
}

func TestCache(size int) {
	myCache := NewLFUCache(size)
	myCache.Update(1, 10)
	myCache.Update(4, 10)
	myCache.Update(1, 20)
	myCache.Update(4, 10)
	myCache.Update(2, 10)
	myCache.Update(3, 10)

	fmt.Println(myCache.Cache, "with count", myCache.Counter)

	test := []int{1, 2, 3, 4, 5}
	fmt.Println(test[:len(test)-1])
}
