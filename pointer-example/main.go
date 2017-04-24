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

func square(x *float64) {
	*x = *x * *x
}

func main() {
	// x := 5
	// zero(x)
	// zeroPoint(&x)

	x := 5.0
	square(&x)
	fmt.Println(x)
}
