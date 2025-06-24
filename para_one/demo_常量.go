package main

import (
	"fmt"
)

// 常量const声明
// 有类型常量，指定常量类型，缺点是不简洁
// 无类型常量，使用const声明常量，不显示指定常量的类型

type NewInt int

// Y 有类型常量
const Y NewInt = 10
const Z int = 20 + int(Y)

// A 无类型常量
const (
	A        = 10
	B        = A
	C NewInt = A
)

func main() {
	fmt.Printf("%d\n", Y)
	fmt.Printf("%d %d %d\n", A, B, C)
	// 基本就是无类型常量用得多
	//http.Server{}
}
