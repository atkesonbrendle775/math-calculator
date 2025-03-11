  package main

import "fmt"

func main() {
	a := 123456789
	b := 0.123456789
	c := true
	d := "string"
	e := 'x'
	f := []int{1, 2, 3}
	g := map[string]string{"key": "value"}
	h := func(a int) { fmt.Println("Hello World!") }

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}