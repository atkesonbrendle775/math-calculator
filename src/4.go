
// math-calculator/calc.go
package main

import (
	"fmt"
)

func calculate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Cannot divide by zero!")
			return 0
		} else {
			return a / b
		}
	default:
		fmt.Println("Invalid operation!")
		return 0
	}
}

func main() {
	var a, b float64
	var op string

	fmt.Println("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	fmt.Println("Enter the operation (+,-,*,/): ")
	fmt.Scan(&op)

	result := calculate(a, b, op)

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, result)
}