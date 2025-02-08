package main

import (
	"fmt"
)

/*
	  --- Parameter VS Argument ---
		Function parameters are declared inside the parentheses while function arguments are passed when the function is called.
		In the given function, a and b are parameters. When we call the function add(10,5), 10 and 5 are arguments.
*/
func add(a int, b int) { // parameter => a,b
	c := a + b
	fmt.Println("Sum:", c)
}

func main() {
	add(10, 5) // argument => 10,5
}
