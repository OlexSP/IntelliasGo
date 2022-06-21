package main

import (
	"fmt"
	"math"
)

var (
	f3 float64 = 10
)

func main() {
	someInt := 10
	fmt.Println("дано", f3, someInt)
	mult := float64(someInt) * f3
	fmt.Println(mult)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxInt)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)

}
