package main

import (
	"fmt"
	"time"
)

// 通过const来定义枚举
// iota 是一个内置的常量生成器，可以用于创建一系列递增的常量值
// iota 无类型枚举常量
const (
	A = iota
	B
	C
	D
)

// Color iota 有类型枚举常量
// 适用于有严格的类型校验
type Color int

const (
	Red Color = iota
	Blue
	Yellow
	Green
)

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Printf("%d, %d, %d, %d\n", A, B, C, D)

	var y int32 = 1
	var x float32 = 3.4
	fmt.Println(y + B)
	fmt.Println(x + B)

	fmt.Println(Red, Blue, Yellow, Green)
	c := Weekday(2)
	fmt.Println(c)
	one := time.Sunday
	fmt.Println(one)
}
