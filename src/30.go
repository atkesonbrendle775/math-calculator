package main

import (
	"fmt"
)

func calculate(a float64, b float64) float64 {
	return a + b
}

func main() {
	result := calculate(2.5, 3.0)
	fmt.Println(result)
}
