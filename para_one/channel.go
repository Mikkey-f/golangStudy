package main

import (
	"fmt"
	"time"
)

// channel 是一种用于协程之间通信和同步的机制
func main() {
	// 初始化channel
	// 没有缓存，又写又读，死锁了
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		data := <-ch
		fmt.Println(data)
	}()
	time.Sleep(time.Second * 1)
}
