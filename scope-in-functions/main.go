package main

import (
	"fmt"
)

/*
---- SCOPES ----
Variables declared outside of a function have global scope.

---- Local ----
Variables declared inside a function have local scope.

---- Block ---
 A block is a group of statements enclosed within curly braces { }.
 -> {}

*/

var (
	a = 30
	b = 40
)

func addNumbers(x, y int) {
	z := x + y

	fmt.Println(z)
}

func main() {

	p := 40
	q := 60
	//local scope
	addNumbers(p, q)

	addNumbers(a, b)

	addNumbers(a, p)

	addNumbers(b, p)
	//global scope
	addNumbers(a, b)
	// z is undefined
	//block scope
	// addNumbers(b,z)

}
