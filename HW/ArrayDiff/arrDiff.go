package main

import "fmt"

func main() {
	//var emptyArr []int
	//a1 := ArrayDiff(emptyArr, []int{1, 2})
	//a2 := ArrayDiff([]int{1, 2, 6, 1}, []int{1})
	a3 := ArrayDiff([]int{1, 2, 2}, []int{2})
	//fmt.Println(a1)
	//fmt.Println(a2)
	fmt.Println(a3)

}

func ArrayDiff(a, b []int) []int {
	for _, v := range b {
		for j := 0; j < len(a); j++ {
			if a[j] == v {
				a = append(a[:j], a[j+1:]...)
				j--
			}
		}
	}
	return a
}
