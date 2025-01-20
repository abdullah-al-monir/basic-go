package main

import (
	"fmt"
)

/*
  ---- Operators ----
	Arithmetic: +, -, *, /
  Comparison: >, <, >=, <=, ==,!=
	Logical: && (AND), || (OR),! (NOT)
	Bitwise: &, |, ^, ~, <<, >>
  Assignment: =, +=, -=, *=, /=, %=
	Ternary: condition? value1 : value2
	Type conversion: type(variable)
	Increment/decrement: ++, --
	Range: for i := range array {...}
	Switch: switch variable { case value1:... case value2:... default:... }
*/

func main() {
	age := 19
	sex := "female"
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age <= 12 {
		fmt.Println("You are a child")
	} else {
		fmt.Println("You are a teenager")
	}

	if age >= 21 && sex == "male" {
		fmt.Println("You are eligible for marriage")
	} else if age >= 18 && sex != "male" {
		fmt.Println("You are eligible for marriage")
	} else {
		fmt.Println("You are not eligible for marriage")
	}

	isSad := false

	if !isSad {
		fmt.Println("I'm feeling sad")
	} else {
		fmt.Println("I'm not feeling sad")
	}

	num := 100

	switch num {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3, 4, 5:
		fmt.Println("Number is 3, 4, or 5")
	default:
		fmt.Println("Number is not 1, 2, or 3, 4, or 5")
	}

}
