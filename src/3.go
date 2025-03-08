package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5
	b := 10

	fmt.Println("The sum of", a, "and", b, "is", add(a, b))
}

func add(a, b int) int {
	return a + b
}
