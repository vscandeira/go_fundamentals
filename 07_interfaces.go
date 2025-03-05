// example code using two geometric figures, a rectangle and a circle, with a method to calculate area initially defined by an interface
package main

import (
	"fmt"
	"math"
)

// GeometricFigure interface
type GeometricFigure interface {
	Area() float64
	Perimeter() float64
}

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

// Circle struct
type Circle struct {
	Radius float64
}

// Area method for the Circle struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// creating a new rectangle
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	// creating a new circle
	circle := Circle{
		Radius: 5,
	}

	// creating a slice of geometric figures
	figures := []GeometricFigure{rectangle, circle}

	// iterating over the slice of geometric figures
	for _, figure := range figures {
		fmt.Println(figure)
		// printing the area of the geometric figure
		fmt.Printf("Area: %f\n", figure.Area())
		fmt.Printf("Perimeter: %f\n", figure.Perimeter())
		fmt.Println()
	}
}
