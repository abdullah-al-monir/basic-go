package main

import (
	"fmt"
)

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {

	sum := num1 + num2

	mul := num1 * num2

	return sum, mul
}

func main() {

	a := 50
	b := 10

	sum := add(a, b)

	getNumbers(a, b)

	sum, mul := getNumbers(a, b)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", mul)
}
