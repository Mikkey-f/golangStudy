package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 1. init()函数 包被导入了自动执行，不能显示的被调用
func init() {
	fmt.Println("init")
}

// 2. main()函数 只能在main包里面
func main() {
	num := sum(1, 2)
	fmt.Println(num)

	data, err := readFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data:", string(data))

	names := []string{"nike", "allen", "jimmy"}
	nameStr := strings.Join(names, ",")
	fmt.Println("nameStr:", nameStr)
}

// 3. 正常函数
// func 名字(参数) 返回值列表 {函数体}
func sum(a, b int) int {
	num := a + b

	fmt.Println("sum:", num)
	// 4. defer 函数 一般释放资源
	// 先执行函数内容。执行完成后再执行defer func() 内容
	defer func() {
		fmt.Println("defer num:", num)
	}()
	return num
}

func readFile(filename string) (string, error) {

	file, err := os.Open(filename)
	defer func() {
		fmt.Println("exec file close")
		_ = file.Close()
	}()
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
