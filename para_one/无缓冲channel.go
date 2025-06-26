package main

//
//import (
//	"fmt"
//	"sync"
//)
//
//// 无缓冲 那就只有一写一读，感觉还是串行
//// 有缓冲 那才是异步并发
//// 死锁问题 select -> 1. 避免无缓冲channel死锁 2. 监听多个channel
//func send(ch chan<- int) {
//	defer close(ch)
//	for i := 0; i < 10; i++ {
//		select {
//		case ch <- i:
//			fmt.Println("发送成功", i)
//		default:
//			fmt.Println("队列空间不足")
//		}
//	}
//	fmt.Println("发送结束")
//	//close(ch)
//}
//
//func receive(ch <-chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for data := range ch {
//		fmt.Println("接收数据", data)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	ch := make(chan int, 10)
//	wg.Add(1)
//	go send(ch)
//	go receive(ch, &wg)
//	wg.Wait()
//}
