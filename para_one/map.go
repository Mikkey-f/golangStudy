package main

import "fmt"

// map(k-v)
// 1. 无序键值对
// 2. key的要求(基础类型用!=判断)
// 3. 零值不可用
// 4. map 引用类型
func main() {
	var names map[int]string
	if names == nil {
		names = make(map[int]string)
	}
	names[1] = "小明"
	//names := map[int]string{
	//	1: "小明",
	//	2: "222",
	//	3: "333",
	//}
	//names[4] = "小蓝"
	//ages := make(map[int]int)
	//ages[1] = 10
	//ages[2] = 20
	//ages[3] = 30
	//fmt.Println(names)
	//fmt.Println(ages)
	//
	//fmt.Println(names[1], ages[1])
	//delete(ages, 2)
	//fmt.Println(ages)
	//for k, v := range names {
	//	fmt.Println(k, v)
	//}
	//fmt.Println(names)
	//updateNames(names)
	//fmt.Println(names)
	//
	//http.Server{}
	ages := make(map[int]int, 10)
	fmt.Println(len(ages))
}

func updateNames(names map[int]string) {
	names[1] = "小王"
}
