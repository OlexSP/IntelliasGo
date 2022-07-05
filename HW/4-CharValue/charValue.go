package main

import (
	"fmt"
)

/*n this Kata, you will be given a string and your task is to return the most
valuable character. The value of a character is the difference between the index
of its last occurrence and the index of its first occurrence. Return the
character that has the highest value. If there is a tie, return the
alphabetically lowest character

func Solve(s string) rune {
    max,r := 0,rune(123)
    for _,c := range s {
        d := LastIndex(s, string(c)) - Index(s, string(c))
        if max < d || max == d && r > c {
            max = d; r = c
        }
    }
    return r
}

*/

func main() {
	string1 := "axyzxyz"
	fmt.Println([]rune(string1))
	fmt.Println(Solve(string1))

}

func Solve(s string) rune {
	var maxDelta, maxInd int
	uniqRuneSlice := []rune(s) //makeUniq(s)
	for i, v := range uniqRuneSlice {
		delta := indexDelta(s, v)
		if delta > maxDelta {
			maxDelta, maxInd = delta, i
		}
		if delta == maxDelta && v < uniqRuneSlice[maxInd] {
			maxInd = i
		}
	}
	return uniqRuneSlice[maxInd]
}

func makeUniq(s string) []rune { // returns unique rune slice
	runeSlice := make([]rune, 0, len(s))
	mapCounter := make(map[rune]int)
	for _, v := range s {
		mapCounter[v]++
	}
	for key := range mapCounter {
		runeSlice = append(runeSlice, key)
	}
	return runeSlice
}

func indexDelta(str string, char rune) int {
	indexSlice := make([]int, 0, len(str))
	for i, v := range str {
		if v == char {
			indexSlice = append(indexSlice, i)
		}
	}
	return indexSlice[len(indexSlice)-1] - indexSlice[0]
}
