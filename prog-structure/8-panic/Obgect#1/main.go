package main

import "fmt"

func main() {
	a := "a"
	b := func(n int) bool {
		fmt.Println(a)
		return false
	}
	d := bar

	foo(b)
	foo(d)

}

func bar(n int) bool {
	return false
}

func foo(f1 func(int) bool) {
	n := 1
	b := f1(n)
	fmt.Println(b)
}
