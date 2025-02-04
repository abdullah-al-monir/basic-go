package main

import (
	"fmt"
)

// Function Expression or Assign function in variable

func sum() {
	add(5, 6)
}

func add(a, b int) {
	c := a + b
	fmt.Println("Old",c)
}
func main() {
	sum()
	add(6,8) // old function
	add := func(a int, b int) {
		c := a + b
		fmt.Println("New",c)
	}
	add(5, 3) // new function
}
