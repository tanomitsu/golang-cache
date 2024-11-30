package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

type Cache3 struct {
	mu    sync.Mutex
	items map[int]int
}

func NewCache3() *Cache3 {
	m := make(map[int]int)
	c := &Cache3{
		items: m,
	}
	return c
}

func (c *Cache3) Set(key int, value int) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

func (c *Cache3) Get(key int) int {
	c.mu.Lock()
	v, ok := c.items[key]
	c.mu.Unlock()

	if ok {
		// キャッシュにkeyが存在した
		return v
	}

	vv, err, _ := group.Do(fmt.Sprintf("cacheGet_%d", key), func() (interface{}, error) {
		value := HeavyGet(key)
		c.Set(key, value)
		return value, nil
	})

	if err != nil {
		panic(err)
	}

	// interface{}型なのでintにキャスト
	return vv.(int)
}
