package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectange Rectangle) float64 {
	return (rectange.Width * 2) + (rectange.Height * 2)
}

func Area(rectange Rectangle) float64 {
	return rectange.Width * rectange.Height
}
