package main

import (
	"fmt"
)

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
	// fmt.Println(y)
}
