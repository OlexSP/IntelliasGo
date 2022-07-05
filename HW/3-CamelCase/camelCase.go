package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text1 := "the-stealth-warrior"
	text2 := "The_Stealth_Warrior"
	fmt.Println(ToCamelCase(text1))
	fmt.Println(ToCamelCase1(text1))
	fmt.Println(ToCamelCase2(text2))
}

func ToCamelCase(s string) string {
	// your code
	runSlice := make([]rune, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '_' || s[i] == '-' {
			i++
			if s[i] > 96 {
				runSlice = append(runSlice, rune(s[i])-32)
				continue
			}
		}
		runSlice = append(runSlice, rune(s[i]))
	}
	return string(runSlice)
}

func ToCamelCase1(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1) // splits into words
	fmt.Println(words)
	for i, w := range words[1:] {
		words[i+1] = strings.Title(w) // every word to Title
	}

	return strings.Join(words, "")
}

func ToCamelCase2(s string) string {
	return regexp.MustCompile("[-_](.)").ReplaceAllStringFunc(s, func(w string) string {
		return strings.ToUpper(w[1:])
	})
}
