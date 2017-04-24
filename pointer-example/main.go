package main

import (
	"fmt"
)

func zero(x int) {
	x = 0
}

func zeroPoint(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	// zero(x)
	zeroPoint(&x)
	fmt.Println(x)
}
