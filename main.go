package main

import (
	"log"
	"time"
)

func main() {
	// pattern1を試す
	mCache := NewCache1()
	log.Println(mCache.Get(3))

	// pattern2を試す
	mCache2 := NewCache2()
	log.Println(mCache2.Get(3)) // 最初はdefaultValueが返る
	time.Sleep(time.Second)
	log.Println(mCache2.Get(3)) // 次は更新されている

	// pattern3を試す
	mCache3 := NewCache3()

	for i := 0; i < 100; i++ {
		go func(i int) {
			// 0から9までのキーをほぼ同時に取得するがそれぞれ一度しかHeavyGetが呼ばれない
			mCache3.Get(i % 10)
		}(i)
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		log.Println(mCache3.Get(i))
	}
}
