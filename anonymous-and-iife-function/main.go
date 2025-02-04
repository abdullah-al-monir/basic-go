package main

import (
	"fmt"
)

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	// anonymous function
	func() {
		fmt.Println("This is an anonymous function.")
	}() // call the anonymous function immediately
	// note: anonymous functions cannot be assigned to variables and cannot be called after they're defined. They can only be called when they're invoked directly.
	// anonymous functions can be useful when you want to perform a single operation without defining a separate function. For example, in a loop or when passing a function as an argument to another function.
	// However, it's important to note that anonymous functions can sometimes make the code harder to read and maintain. It's generally recommended to define named functions whenever possible.
	// Anonymous functions can also be used to create closures, which are functions that remember and access variables from their surrounding scope. This is a powerful feature in functional programming.

	/*
	  Immediately Invoked Functions Expressions
	*/
  add(5,8)
}
