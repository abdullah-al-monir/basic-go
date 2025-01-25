package main

import (
	"fmt"
)

var a = 10

func main() {
	age := 30

	if age > 18 {
		a := 45
		fmt.Println(a) // Output a = 45
	}
	fmt.Println(a) // Output a = 10
}
