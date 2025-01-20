package main

import (
	"fmt"
)

/*
 ---- Variable Declaration ----
 var x, y, z int
 p := 10 // := is shorthand for declaring and initializing a variable

 ---- Var ----
 var a, b, c = 1, 2, 3 // multiple variable declaration and initialization at once

 ---- Constants ----
 const Pi = 3.14159 // The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

*/

/*
 ---- Data Types ----
  int
  float32
	float64
	string
	bool
*/

func main() {
	//int
	var x int = 10
	const Pi = 3.14159
	// We also can declare variables using ':=' syntax
	p := 15

	// We can reassign variables but type should be same
	p = 20

	//float
	y := 3.14159

	//string
	z := "We don't talk anymore..."

	//boolean
	isTrue := true

	fmt.Println("Pie -", Pi)
	fmt.Println("x =", x, "& p =", p)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
	fmt.Println("isTrue =", isTrue)

}
