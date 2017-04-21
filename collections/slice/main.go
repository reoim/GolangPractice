package main

import (
	"fmt"
)

func main() {
	sliceA := make([]int, 0, 3)

	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	sliceB = append(sliceB, sliceC...)
	// sliceB = append(sliceB, 4,5,6)
	fmt.Println(sliceB)
}
