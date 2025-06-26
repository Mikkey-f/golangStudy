package main

import (
	"errors"
	"fmt"
)

// ErrNotFound Error格式
var ErrNotFound = errors.New("not found")

// 常见的error创建的两种方式
func do() error {
	//return errors.New("test error")
	return fmt.Errorf("test error id = %d", 1)
}

// Find error一般放在返回值的最右侧
func Find(id int) (int, NetError) {
	if id == 1 {
		return 0, &MyError{
			Code: -2001,
			Msg:  "test error",
		}
	}
	return 1, nil
}
func main() {
	data, err := Find(1)
	if err != nil {
		// 在业务逻辑中这种是正常的, 哨兵机制
		//if err == ErrNotFound {
		//	// SET
		//	data = 2
		//} else {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println(err.(*MyError).Code)
		//return
		// ok 为是否转换成功
		if e, ok := err.(*MyError); ok && e.IsLink() {
			// 重试
			fmt.Println("进行重试")
		} else {
			return
		}
	}
	fmt.Println(data)
}

// 自定义error
type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return "my error"
}

func (e *MyError) IsLink() bool {
	return true
}

// 行为特征策略
type NetError interface {
	error
	IsLink() bool
}
