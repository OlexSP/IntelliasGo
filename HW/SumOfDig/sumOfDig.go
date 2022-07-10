package main

import "fmt"

func main() {
	fmt.Println(DigitalRoot(942))

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
