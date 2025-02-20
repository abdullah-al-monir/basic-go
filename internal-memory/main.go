package main

import (
	"fmt"
)

// Data Segment (Initialized global variable)
var a = 10

// init() is executed before main(), part of the Code Segment
func init() {
	fmt.Println("Initializing...")
}

// add() function resides in the Code Segment
func add(x int, y int) {
	// Stack Segment (Local variables: x, y, z)
	z := x + y // z is pushed onto the stack
	fmt.Println("Sum:", z)
} // z is popped from the stack when function exits

func main() {
	// Function calls use the Stack Segment
	add(5, 4) // x=5, y=4, z=9 stored in stack
	add(a, 5) // a is in Data Segment, y=5 in stack

	// Heap Segment (Demonstrating heap allocation)
	p := new(int)                            // p is a pointer allocated on heap
	*p = 20                                  // Assigning value to heap memory
	fmt.Println("Heap allocated value:", *p) // Accessing heap memory
} // p is deallocated by Go’s garbage collector

/*
  2 Phases:
	 1. Compilation Phase
	 2. Execution Phase
  
	 *** Compile Phase ***

	 *** Code Segment ***
	 a = 10
	 add = func (){...}
	 main = func() {...}
	 init = func() {...}



	go run main.go => compile it => main => ./main
	go build main.go => compile it => main

*/

/*
 *** Binary File ***

  main

	0011010101101
	0011010101101
	0011010101101
	0011010101101
	0011010101101
	0011010101101




*/
