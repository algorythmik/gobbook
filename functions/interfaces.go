package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	return math.Abs(r.x1-r.x2) * math.Abs(r.y1-r.y2)

}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	area := 0.0
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func totalArea(shapes ...Shape) float64 {
	area := 0.0
	for _, s := range shapes {
		area += s.area()

	}
	return area
}

func main() {
	c := Circle{0, 0, 10}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(totalArea(&c, &r))
	m := []Shape{&c, &r}
	m2 := MultiShape{m}
	fmt.Println(m2.area())
}
