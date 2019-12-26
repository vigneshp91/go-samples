package main

import "fmt"

type calculations interface {
	arear() float64
}

type circle struct {
	radius float64
}

func (c circle) arear() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	fmt.Print("hello")
	c := circle{4}
	fmt.Println(c.arear())

}
