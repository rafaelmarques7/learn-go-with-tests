package smi

import "math"

// The Shape interface is valid for any struct that contains an Area() method
// Thus, it will be valid both for Rectangle and for Circle types
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}
