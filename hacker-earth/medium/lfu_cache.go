package medium

import (
	"container/list"
	"fmt"
)

type LFUCache struct {
	Cache       map[int]int
	Counter     map[int]int64
	Size        int
	CurrentSize int
	LFU         []*list.List
	MinCounter  int64
}

func NewLFUCache(size int) *LFUCache {
	init := list.New()
	return &LFUCache{
		Cache:       make(map[int]int),
		Counter:     make(map[int]int64),
		Size:        size,
		CurrentSize: 0,
		LFU:         []*list.List{init},
		MinCounter:  0,
	}
}

func (c *LFUCache) UsedSpace() int {
	return len(c.Cache)
}

func (c *LFUCache) IsFull() bool {
	return c.CurrentSize >= c.Size
}

func (c *LFUCache) Get(key int) int {
	if val, ok := c.Cache[key]; ok {
		_count := c.Counter[key]

		// * Get key in current counter
		l := c.LFU[_count]
		el := l.Front()
		for {
			if el == nil {
				break
			}
			if el.Value.(int) == key {
				// * found the key in freqs, remove it and move to next freqs
				l.Remove(el)
				break
			} else {
				el = el.Next()
			}
		}

		// * Increase counter of key
		if _count == c.MinCounter && l.Len() == 0 {
			c.MinCounter++
		}

		_count++
		if len(c.LFU) == int(_count) {
			c.LFU = append(c.LFU, list.New())
		}
		c.LFU[_count].PushBack(key)
		c.Counter[key] = _count
		return val
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if c.Get(key) != -1 {
		// * in case key is already exist, just update and increase counter
		// * The counter and LFU is mutated in Get()
		c.Cache[key] = value
		return
	}

	if c.IsFull() {
		// * in case cache is full, invalidate min freq cache before add new
		minEl := c.LFU[c.MinCounter].Front()
		minKey := c.LFU[c.MinCounter].Remove(minEl)
		c.Delete(minKey.(int))
	}

	// * Put new key/value into Cache
	c.Cache[key] = value
	c.Counter[key] = 0
	c.MinCounter = 0
	c.CurrentSize++
	c.LFU[0].PushBack(key)
}

// clear cache also clear counter of cached key
func (c *LFUCache) Delete(key int) {
	delete(c.Cache, key)
	delete(c.Counter, key)
	c.CurrentSize--
}

func TestCache() {
	// * INPUT: [[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	// * OUTPUT: [null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]
	CacheOpts := [][]int{
		{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26},
	}
	myCache := NewLFUCache(10)

	for _, opt := range CacheOpts {
		// fmt.Printf("action %v with cache state %v - %v\n", opt, myCache.Cache, myCache.LFU)
		if len(opt) == 1 {
			myCache.Get(int(opt[0]))
			// fmt.Printf("%d ", val)
		} else {
			myCache.Put(int(opt[0]), opt[1])
			// fmt.Printf("null ")
		}
	}
	fmt.Println("\n", myCache.Cache, "with count", myCache.Counter)
}
