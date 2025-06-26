package main

import (
	"fmt"
	"net/http"
)

// 方法是指在类型上定义的函数，它们可以像普通的函数一样被调用，但需要使用接收者来访问其定义的结构体
// 接收者receiver
func main() {
	s := &school{
		name: "school",
		addr: "校区",
	}
	name := s.getName()
	fmt.Println("name =", name)
	http.Server{}
}

type school struct {
	name string
	addr string
}

// s *school指明s是school对象的一个引用
func (s *school) getName() string {
	return s.name
}

// 上述方法其实就是在编译期间转换成了下列函数
//func getName(s * school) string {
//	return s.name
//}

// 注意
// 1. (函数名) 首字母大写叫做可导出的方法(即其他包可以调用)
// 2. 不能为基础类型添加方法, 不能跨包为自定义类型添加方法
// 3. 方法的本质就是函数和数据类型组合在一起
