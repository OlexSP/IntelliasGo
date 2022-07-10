package main

import "fmt"

func main() {
	fmt.Println(DigitalRoot(942))
	fmt.Println(DigitalRootRecursion(9425))
	fmt.Println(DigitalRoot2(942))

}

func DigitalRoot(n int) int {
	sum := 0
L:
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	if sum > 9 {
		n, sum = sum, 0
		goto L
	}

	return sum
}

func DigitalRootRecursion(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	if sum > 9 {
		return DigitalRootRecursion(sum)
	}

	return sum
}
func DigitalRoot2(n int) int {
	return (n-1)%9 + 1
}
