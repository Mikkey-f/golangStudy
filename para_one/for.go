package main

import (
	"fmt"
	"strings"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	// while 循环
	//i := 0
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}
	// do while 循环
	//i := 0
	//for {
	//	i++
	//	fmt.Println(i)
	//	if i == 10 {
	//		break
	//	}
	//}
	//names := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	//for index, name := range names {
	//	fmt.Println(index, name)
	//}

	name := "allen"
	for _, v := range name {
		fmt.Println(string(v))
	}

	ages := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range ages {
		fmt.Println(k, v)
	}
	strings.Split()
}
