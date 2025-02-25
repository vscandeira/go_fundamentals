package main

import "fmt"

func main() {
	var name string = "John"
	var lastName = "Doe"
	fmt.Println(name + " " + lastName)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of lastName: %T\n\n", lastName)

	var a int = 1
	var b = 2
	fmt.Println(a + b)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n\n", b)

	var c float32 = 1.5
	var d = 2.5
	fmt.Println(float64(c) + d)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n\n", d)

	var e bool = true
	var f = false
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("Type of e: %T\n", e)
	fmt.Printf("Type of f: %T\n\n", f)

	var g int = 10
	var h = 20
	fmt.Println(g == h)
	fmt.Println(g != h)
	fmt.Println(g > h)
	fmt.Println(g < h)
	fmt.Println(g >= h)
	fmt.Println(g <= h)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Printf("Type of g: %T\n", g)
	fmt.Printf("Type of h: %T\n\n", h)
}
