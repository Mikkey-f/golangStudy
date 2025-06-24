package main

import (
	"fmt"
)

var name string
var age int

// 使用规范
/**
使用规范：
局部变量立刻赋值使用短变量:=
用var声明一类的变量
*/

// 全局变量
var (
	x = int64(1)
	y = float32(1.2)
	z = "1"
)

func main() {
	fmt.Println("Hello World")
	name = "mikkeyf"
	age = 10
	fmt.Printf("name = %s \n age = %d\n", name, age)
	fmt.Printf("x = %d y = %f z = %s", x, y, z)
	add()
}

// 局部变量
func add() {
	// 短变量声明
	a := 1
	b := 2
	fmt.Printf("a = %d , b = %d, sum = %d\n", a, b, a+b)

	// 长变量声明
	var x, y int
	x = 1
	y = 2
	fmt.Printf("x = %d , y = %d sum = %d\n", x, y, x+y)

	var (
		m = int64(1)
		n = int64(2)
	)
	fmt.Printf("m = %d , n = %d , sum = %d \n", m, n, m+n)
}
