package main

import "fmt"

func main() {

	table := MultiplicationTable(3)
	fmt.Println(table)
}

func MultiplicationTable(size int) [][]int {
	// Implement me! :)
	area := make([][]int, size)
	for i := range area {
		area[i] = make([]int, size)
		for j := range area[i] {
			area[i][j] = (i + 1) * (j + 1)
		}
	}

	return area
}
