package main

import "math/rand"

func main() {

}

func describeSlice(n int) {
	slices := generateSlice(n, elemsInSlice, maxNumber)

}

func generateSlice(n, elemsInSlice, maxNumber int) [][]int {
	slices := make([][]int, n)
	for i := 0; i < n; i++ {
		slice := make([]int, elemsInSlice)

		for j := 0; j < elemsInSlice; j++ {
			slice[j] = rand.Intn(maxNumber)
		}
	}
}

func describeSice(s []int) (sum, min, max, mode int) {
	min = maxNumber
	entriesCount := make(map[int]int)

	for _, elem := range slice {
		sum += elem

		if elem < min {
			min = elem
		}

		if elem > max {
			max = elem
		}
	}
}
