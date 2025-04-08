package main

import "fmt"

func main() {
  // Define some constants and variables
  PI := 3.14159265358979323846f
  RADIUS := 5

  // Calculate the area of a circle using the formula π * R²
  AREA := PI * RADIUS * RADIUS

  // Print the result
  fmt.Printf("The area of the circle is: %.1f\n", AREA)
}
