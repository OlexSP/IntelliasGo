package main

import (
	"fmt"
	"math"
)

const Pi = math.Pi

func main() {
	fmt.Println("pi ", Pi)
	fmt.Printf("%T: %[1]v ", Pi)

	const (
		a = 1
		b = "hello"
		c = true
		e = 1234.5678
	)

}
