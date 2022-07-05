package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "man i need a taxi up to ubdu"

	fmt.Println(High(text))
}

func High(s string) string {
	// your code here
	var maxScore, maxIndex int
	stringSlice := strings.Split(s, " ")
	for i, v := range stringSlice {
		if score := scoreWord(v); score > maxScore {
			maxScore = score
			maxIndex = i
		}
	}
	return stringSlice[maxIndex]
}

func scoreWord(word string) int {
	sum := 0
	for _, v := range word {
		sum += int(v - 96)
	}
	return sum
}
