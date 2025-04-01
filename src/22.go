package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    _, err := fmt.Scan(&a, &b)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := a + b
    fmt.Println("The sum of", a, "and", b, "is:", result)

    var expression string = "expression"
    exprFunc(expression)
}

func exprFunc(expr string) {
    // Add your Go code for the expression here
}
