package main

import (
	"sync"
)

type Cache1 struct {
	mu    sync.Mutex
	items map[int]int
}

func NewCache1() *Cache1 {
	m := make(map[int]int)
	c := &Cache1{
		items: m,
	}
	return c
}

func (c *Cache1) Set(key int, value int) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

func (c *Cache1) Get(key int) int {
	c.mu.Lock()
	v, ok := c.items[key]
	c.mu.Unlock()

	if ok {
		// キャッシュにkeyが存在した
		return v
	}

	v = HeavyGet(key)

	c.Set(key, v)

	return v
}
