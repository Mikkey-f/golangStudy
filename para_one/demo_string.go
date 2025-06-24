package main

import (
	"fmt"
)

// string定义
// 1. 数据不可变
// 2. 默认值可用
// 3. 支持字符串拼接（+， +=）
// 4. 支持比较字符串（!=, ==)
// 5. 支持多行定义
func main() {
	name := "lucky"
	nameCopy := []byte(name)
	nameCopy[1] = 'l'

	fmt.Println(name, string(nameCopy))
	var school string
	fmt.Println("炸" + school)

	var newName string
	newName += name + " day"
	fmt.Println(newName)
	newName += "s"
	fmt.Println(newName)

	if name == newName {
		fmt.Println("OK")
	} else {
		fmt.Println("FAIL")
	}

	text := `ABCD
				E
F
GGG`
	fmt.Println(text)
	//strings.Split()
}
