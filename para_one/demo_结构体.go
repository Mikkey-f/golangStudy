package main

import (
	"fmt"
	"net/http"
)

type student struct {
	id    int
	name  string
	age   int
	score float64
}

func main() {
	// go推荐的两种方式
	// 1. 使用var
	//var s student
	//s.name = "mikkeyf"
	//s.age = 20
	//s.score = 100

	// 2. 使用student间接赋值
	s := student{
		id:    1,
		name:  "<UNK>",
		age:   20,
		score: 100.0,
	}
	p := &s
	fmt.Println(s.name, s.age, s.score)
	fmt.Println(p)
	p.age = 18
	fmt.Println(p)
	http.Server{}
}
