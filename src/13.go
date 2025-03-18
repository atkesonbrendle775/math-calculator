package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Ask the user to enter two numbers
	var num1, num2 float64
	fmt.Println("Enter the first number: ")
	fmt.Scanf("%f", &num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanf("%f", &num2)

	// Ask the user to choose an operation (+, -, *, /)
	var op string
	fmt.Println("Choose an operation (+, -, *, /): ")
	fmt.Scanf("%s", &op)

	// Perform the calculation and print the result
	switch op {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero.")
		} else {
			fmt.Println(num1 / num2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
