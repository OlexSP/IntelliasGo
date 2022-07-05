package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	mapCounter := make(map[int]int)

	for _, v := range arr {
		mapCounter[v]++
		if mapCounter[v] == 1 {
			result = append(result, v)
		}
	}

	fmt.Printf("Duplicates have been removed: %v\n", result)

}
