package main

import (
	"fmt"
	"math"
)

func main() {
	c := circle{5}
	s := square{10}
	fmt.Println("circle area", c.area())
	fmt.Println("square area", s.area())
}

type circle struct {
	radius float64
}
type square struct {
	length float64
}
type area interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s *square) area() float64 {
	return s.length * s.length
}
