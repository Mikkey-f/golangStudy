package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	//num := rand.Intn(100)
	//
	//if num%2 == 0 {
	//	fmt.Println("num%2 yes", num)
	//} else if num%3 == 0 {
	//	fmt.Println("num%3 yes", num)
	//} else {
	//	fmt.Println("no", num)
	//}
	if num := rand.Intn(10); num%2 == 0 {
		fmt.Println("num%2 yes", num)
	} else {
		fmt.Println(num)
	}

	//http.Server{}

}
