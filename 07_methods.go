// example code using a rectangle struct with a method to calculate the area
package main

import (
	"fmt"
)

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for the Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// creating a new rectangle
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	// printing the rectangle
	fmt.Println("Rectangle:", rectangle)

	// calculating the area of the rectangle
	area := rectangle.Area()

	// printing the area of the rectangle
	fmt.Println("Area:", area)

	// calculating the perimeter of the rectangle
	fmt.Println("Perimeter:", rectangle.Perimeter())

	r2 := Rectangle{20, 40}
	fmt.Println("\nRectangle:", r2)
	fmt.Println("Area:", r2.Area())
	fmt.Println("Perimeter:", r2.Perimeter())
}
