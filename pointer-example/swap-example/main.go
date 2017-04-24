package main

import (
	"fmt"
)

func swap(xPtr, yPtr *int) {
	*xPtr, *yPtr = *yPtr, *xPtr

	//or
	// temp := *xPtr
	// *xPtr = *yPtr
	// *yPtr = temp

}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x is ", x)
	fmt.Println("y is ", y)

}
