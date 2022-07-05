package main

import "fmt"

func main() {
	s := "f1"
	f1 := func(s string) {
		fmt.Println(s)
	}
	f1(s)
	s = "f2"
	f1(s)

}

func foo(n int) {
	defer func() {
		fmt.Println("always be done")
	}()

	if n == 0 {
		return
	}
	if n == 1 {
		return
	}
	if n == 2 {
		return
	}
	if n == 3 {
		return
	}
}
