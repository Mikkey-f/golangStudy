package main

import (
	"fmt"
	"sync"
	"time"
)

// 原生并发
// 使用协程来实现并发的

var mans = map[int64]*man{
	1: {
		id:      1,
		name:    "nike",
		age:     20,
		address: "南区",
	},
	2: {
		id:      2,
		name:    "mikkey",
		age:     30,
		address: "北区",
	},
}

type man struct {
	id      int64
	name    string
	age     int64
	address string
}

func getStudentById(id int64) (*man, error) {
	// 模拟网络请求
	time.Sleep(time.Millisecond * 100)
	return mans[id], nil
}

func main() {
	start := time.Now()
	ids := []int64{1, 2}
	idToAgeMapper := make(map[int64]int64)

	var (
		eg    sync.WaitGroup
		mutex sync.Mutex
	)

	for _, id := range ids {
		eg.Add(1)
		go func(id int64) {
			defer func() {
				eg.Done()
				mutex.Unlock()
			}()
			s, err := getStudentById(id)
			if err != nil {
				return
			}
			mutex.Lock()
			idToAgeMapper[s.id] = s.age
			fmt.Printf("id = [%d], s = [%+v] \n\n", id, s)
		}(id)
	}

	eg.Wait()

	end := time.Since(start)
	fmt.Println(end)
	time.Sleep(time.Minute * 1)
}
