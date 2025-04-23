package shapes

import (
	"math"
)

// Rectangles
type Rectangle struct {
	Width float64
	Height float64
}


func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2*(rectangle.Width + rectangle.Height)

}

// Circles
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64{
	return c.Radius * c.Radius * math.Pi
}

// Triangle
type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5 
}

