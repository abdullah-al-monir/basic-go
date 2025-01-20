package main

import (
	"fmt"
)

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Sum:", sum)
}

func subtract(num1 int, num2 int) {
	diff := num1 - num2
	fmt.Println("Difference:", diff)
}

func multiply(num1 int, num2 int) {
	product := num1 * num2
	fmt.Println("Product:", product)
}

func divide(num1 int, num2 int) {
	if num2 != 0 {
		quotient := num1 / num2
		fmt.Println("Quotient:", quotient)
	} else {
		fmt.Println("Error: Division by zero")
	}
}

func modulus(num1 int, num2 int) {
	if num2 != 0 {
		remainder := num1 % num2
		fmt.Println("Remainder:", remainder)
	} else {
		fmt.Println("Error: Division by zero")
	}
}

func power(base int, exponent int) {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	fmt.Println("Power:", result)
}

func main() {

	p := 70
	q := 10

	add(p, q)
	subtract(p, q)
	multiply(p, q)
	divide(p, q)
	modulus(12, 5)
	power(2, 4)

}
