package main

import "sync"

const defaultValue = 0

type Cache2 struct {
	mu    sync.Mutex
	items map[int]int
}

func NewCache2() *Cache2 {
	m := make(map[int]int)
	c := &Cache2{
		items: m,
	}
	return c
}

func (c *Cache2) Set(key int, value int) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

func (c *Cache2) Get(key int) int {
	c.mu.Lock()
	v, ok := c.items[key]
	c.mu.Unlock()

	if ok {
		// キャッシュにkeyが存在した
		return v
	}

	go func() {
		// キャッシュの更新処理を非同期に走らせる
		v := HeavyGet(key)

		c.Set(key, v)
	}()

	// 規定値をとりあえず返す
	return defaultValue

}
