package main

import (
	"fmt"
)

// 1. Standard (Named) Function
func add(a int, b int) int {
	return a + b
}

// 2. Anonymous Function
var anonymousFunc = func(a, b int) int {
	return a * b
}

// 3. Function Literal (Assigned to variable)
var subtract = func(a, b int) int {
	return a - b
}

// 4. Slice Function
func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// 5. Map Function
func printMap(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 6. Channel Function
func sendData(ch chan int) {
	ch <- 42
}

// 7. Variadic Function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 8. Closure Function
func incrementer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 9. Defer Function
func deferExample() {
	defer fmt.Println("Deferred execution")
	fmt.Println("Normal execution")
}

// 10. Receiver Function (Method)
type Person struct {
	name string
}

func (p Person) greet() {
	fmt.Println("Hello", p.name)
}

// 11. Higher-Order Function
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 12. Callback Function
func doWork(callback func(string)) {
	callback("Task completed")
}

// 13. Init Function
func init() {
	fmt.Println("Initializing application...")
  
}
func main() {
	// Standard function
	fmt.Println("Addition:", add(10, 5))

	// Anonymous function
	fmt.Println("Multiplication:", anonymousFunc(3, 4))

	// Function literal
	fmt.Println("Subtraction:", subtract(10, 5))

	// Slice function
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum of slice:", sumSlice(slice))

	// Map function
	m := map[string]int{"apple": 5, "banana": 10}
	printMap(m)

	// Channel function
	ch := make(chan int)
	go sendData(ch)
	fmt.Println("Received from channel:", <-ch)

	// Variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// Closure function
	inc := incrementer()
	fmt.Println("Increment:", inc())
	fmt.Println("Increment:", inc())

	// Defer function
	deferExample()

	// Receiver function
	p := Person{"Alice"}
	p.greet()

	// Higher-order function
	result := operate(10, 5, add)
	fmt.Println("Higher-order function result:", result)

	// Callback function
	doWork(func(msg string) {
		fmt.Println("Callback message:", msg)
	})

	// Immediately Invoked Function Expression (IIFE)
	func() {
		fmt.Println("This is an IIFE in Go")
	}()
}
