package main

import (
	"fmt"
)

/*
***** First order Functions VS Higher Order Functions *****

1.First order functions: Functions that take other functions as arguments or return them as results.
 i. Standard functions or named functions
 ii. Anonymous functions
 iii. Function literals (assigned to variables)
 iv. Closures (functions that remember and access variables from their surrounding scope)
 v. Immediately Invoked Function Expressions (IIFEs)

--------------------------------------------------

2.Higher order functions: Functions that take one or more functions as arguments, or return a function as a result.

3. Callback functions: Functions that take one or more functions as arguments

4. First Class Citizen => (Assignable data in variables to variables)
*/

// Higher order function example
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func call() func(p int, q int) {
	return add
}

func add(x int, y int) {
	z := x + y
	fmt.Println("Sum:", z)
}
func main() {
	processOperation(5, 7, add)
	sum := call()
	sum(4, 6)
	a := true
	b := false
	if a {
		fmt.Println("True")
	} else if b {
		fmt.Println("False")
	} else {
		fmt.Println("Neither True nor False")
	}
}
