package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	sideLen float64
}

func (s square) area() float64 {
	return s.sideLen * s.sideLen
}

type shape interface {
	area() float64
}

type outPrinter struct{}

func (op outPrinter) toText(s shape) string {
	return fmt.Sprintf("the area is: %5.2f", s.area())
}

func main() {
	c := circle{radius: 5}

	s := square{sideLen: 2}

	out := outPrinter{}

	fmt.Println(out.toText(c))
	fmt.Println(c)
	fmt.Println(out.toText(s))

}
