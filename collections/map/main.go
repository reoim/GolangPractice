package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[111] = "Reo"
	m[222] = "Hyeran"
	m[223] = "Haemin"
	str := m[111]
	m[111] = "Reo Lee"
	fmt.Println(str)
	fmt.Println(m[111])

	// check map key
	val, exists := m[111]
	if !exists {
		fmt.Println("No 111 key")
	}
	fmt.Println("check map key and return value: " + val)

	// for range loop
	for key, val := range m {
		fmt.Println(key, val)
	}
}
