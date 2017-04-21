package vis

import (
	"fmt"
)

func PrintVar() {
	fmt.Println(MyFirstName)
	// myLastName is not exported outside package but it is accessabile inside package
	fmt.Println(myLastName)
}
