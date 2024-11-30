package main

import (
	"log"
	"time"
)

// 実際にはデータベースへの問い合わせが発生する
func HeavyGet(key int) int {
	log.Printf("call HeavyGet(%d)\n", key)
	time.Sleep(time.Second)
	return key * 2
}
