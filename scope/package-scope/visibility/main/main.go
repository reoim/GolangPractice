package main

import (
	"fmt"

	"github.com/reoim/GolangPractice/scope/package-scope/visibility/vis"
)

func main() {
	fmt.Println(vis.MyFirstName)
	// fmt.Println(vis.myLastName)
	// cannot access to vis.myLastName because it is not exported to outside package
	vis.PrintVar()
}
