package smi

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
