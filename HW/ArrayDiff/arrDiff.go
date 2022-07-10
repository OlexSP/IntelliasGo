package main

import "fmt"

func main() {
	var emptyArr []int
	a1 := ArrayDiff(emptyArr, []int{1, 2})
	a2 := ArrayDiff([]int{1, 2, 6, 8}, []int{1, 2})
	fmt.Println(a1)
	fmt.Println(a2)

}

func ArrayDiff(a, b []int) []int {
	for _, vb := range b {
		for i, va := range a {
			if va == vb {
				a = append(a[:i], a[i+1:]...)
				//a = a[:i+copy(a[i:], a[i+1:])]
			}
		}
	}
	return a
}
