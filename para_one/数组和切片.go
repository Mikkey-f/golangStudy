package main

import (
	"fmt"
)

// go中数组是一个固定长度的数据结构
// 1. 长度在创建之后无法改变
// 2. 同一类型
// 3. 值传递

// 切片(slice)是一种动态数组, 可以按需自动增长和缩小。 对底层数组的一种描述
func main() {
	// 短变量 var
	intList := [5]int{1, 2, 3, 4, 5}
	// 值传递->值拷贝->性能损耗->怎么处理->切片
	updateArray(intList)
	fmt.Println(intList)

	//var 关键字初始化
	//var slice []int
	//slice = append(slice, 1)
	//slice = append(slice, 2)
	//slice = append(slice, 3)
	//slice = append(slice, 4)
	//slice = append(slice, 5)
	//slice = append(slice, 6)
	//fmt.Println(slice)

	// 使用make关键字初始化  指针，len   cap >= len
	slice := make([]int, 6) // 长度6, 容量6
	slice[0] = 1
	fmt.Println(slice, len(slice), cap(slice))

	slice1 := make([]int, 6, 12)
	fmt.Println(slice1, len(slice1), cap(slice1))

	// 切片化: 语法u[low: high] 创建对已经存在数组进行操作的切片
	intList1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(intList1, len(intList1), cap(intList1))
	newSlice := intList1[1:3] // [1, 3)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice[0] = 20
	fmt.Println(intList1, newSlice)

	// slice 扩容
	//var slice2 []int
	//slice2 = append(slice2, 1)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //1 1
	//slice2 = append(slice2, 2)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //2 2
	//slice2 = append(slice2, 2)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //3 4
	//slice2 = append(slice2, 2)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //4 4
	//slice2 = append(slice2, 1)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //5 8
	//slice2 = append(slice2, 2)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //6 8
	//slice2 = append(slice2, 2)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //7 8
	//slice2 = append(slice2, 2)
	//fmt.Println(slice2, len(slice2), cap(slice2)) //8 8
	var slice3 []int
	slice3 = append(slice3, 1, 2, 3, 4)
	fmt.Println(slice3, len(slice3), cap(slice3))

	newSlice2 := slice3[1:3]
	fmt.Println(newSlice2, len(newSlice2), cap(newSlice2))
	newSlice2[0] = 20
	fmt.Println(slice3, newSlice2)

	// 扩容后 切片和新数组会解绑
	slice3 = append(slice3, 5)
	newSlice2[0] = 7
	fmt.Println(slice3, newSlice2)

	// go建议显示设置一下cap容量
	//http.Server{}
}

func updateArray(arr [5]int) {
	arr[0] = 11
	fmt.Println(arr)
}
