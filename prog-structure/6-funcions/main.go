package main

import "fmt"

// operates on copies of a,b. do not provide any changes
func swap(a, b int) {
	a, b = b, a
}

// uses pointers to make changes
func swapRight(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 17
	b := 99
	fmt.Printf("a = %d, b = %d\n", a, b)

	swap(a, b) // operates on copies of a,b
	fmt.Println("func swap(a, b int) {a, b = b, a}")
	fmt.Printf("a = %d, b = %d\n", a, b) // a,b has not been changed

	swapRight(&a, &b) // uses pointers
	fmt.Println("func swapRight(a, b *int) {*a, *b = *b, *a}")
	fmt.Printf("a = %d, b = %d\n", a, b) // a,b has been changed

	/*testSlice := []int{1, 2, 3}
	fmt.Println(testSlice...)*/ // Println has to accept []interface{},  []int - not meet func recs

}

func getMinMax(s []int) (min int, max int) { // сігтнатура / вхідні.вихідні параметри
	if len(s) == 0 {
		return 0, 0
	}
	min = s[0]
	max = s[0]

	for _, v := range s {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

/*  refactoring of a bad func "fast return"
func canBePresident(age, yearsInUkraine int, citizenship string) bool {
	var result bool

	if age < 35 {
		result = false
	} else {
		if yearsInUkraine < 10 {
			result = false
		} else {
			if citizenship != "Ukraine" {
				result = false
			} else {
				result = true
			}
		}
	}

	return result
}
*/
func canBePresident(age, yearsInUkraine int, citizenship string) bool {
	if age < 35 { // reading func logics is much better than canBePresident2
		return false // return faster
	}
	if yearsInUkraine < 10 {
		return false
	}
	if citizenship != "Ukraine" {
		return false
	}
	return true
}

func canBePresident2(age, yearsInUkraine int, citizenship string) bool {
	return age >= 35 && yearsInUkraine >= 10 && citizenship == "Ukraine"
} // it is recommended put in the first place the most complicated func for faster output
