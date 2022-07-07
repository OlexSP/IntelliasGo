package main

import "fmt"

func main() {
	//fmt.Println("main(): start")
	someFunc1()
	//fmt.Println("main(): end")
}

func someFunc1() {
	/*defer func() {
		recover()
		fmt.Println("defer 1")
	}()

	fmt.Println("someFunc1(): start")*/
	n := someFunc2()
	fmt.Println(n)
	/*fmt.Println("someFunc1(): end")
	defer func() {
		fmt.Println("defer 2")
	}()*/
}

func someFunc2() (n int) {
	defer func() {
		recover()
		//fmt.Println("defer 3")
		n = 2
	}()

	//fmt.Println("someFunc2(): start")
	panic("help me!")
	//fmt.Println("someFunc2(): end")

	return 1
}
