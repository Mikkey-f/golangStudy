package main

import (
	"fmt"
	"runtime"
	"time"
)

// panic 异常处理机制，用于表示程序发生了无法处理的错误或异常情况
func main() {
	// 触发panic
	//panic("panic test")
	//var name []string
	//fmt.Println(name[0])
	go func() {
		// panic 恢复
		defer func() {
			if err := recover(); err != nil {
				// panic 处理
				const size = 64 << 10
				buf := make([]byte, size)
				// 打印当前协程信息
				buf = buf[:runtime.Stack(buf, false)]
				fmt.Println(err, string(buf))
			}
		}()
		var names = []string{"1"}
		fmt.Println(names[0])
	}()

	time.Sleep(10 * time.Second)
}

// 常见应用场景
// 1. 初始化配置（redis，mysql连接等等)
// 2. 处理不应该发生的异常情况 （比如计算面积传入复数)
// 3. 确保接口和抽象类型的实现正确性
