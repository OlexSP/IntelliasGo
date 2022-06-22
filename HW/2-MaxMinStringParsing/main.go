package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	var intSlice []int
	sign, intDig := 1, 0
	for _, v := range input {
		switch v {
		case ' ':
			continue
		case '-':
			sign = -1
		default:
			intDig, _ = strconv.Atoi(string(v))
			intSlice = append(intSlice, intDig*sign)
			sign = 1
		}
	}
	sort.Ints(intSlice)
	result = fmt.Sprintf("%v %v\n", intSlice[len(intSlice)-1], intSlice[0])
	fmt.Println(result)
}
