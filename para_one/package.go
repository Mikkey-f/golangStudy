package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	err := errors.New("Hello Error")
	fmt.Println(err)
	
}
