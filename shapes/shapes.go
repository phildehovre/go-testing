package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (c Circle) Perimeter() float64 {
	return (math.Pi * c.Radius) * 2
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Base * 3
}
