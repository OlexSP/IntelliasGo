package main

import "fmt"

func main() {

	var m1 map[string]int // nil map
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 0,
	}
	fmt.Println("map m2:", m2)
	fmt.Println("значення м2 по ключу а:", m2["a"])
	fmt.Println("значення м1 по ключу а:", m1["a"])
	dValue, ok := m2["d"]
	fmt.Println("key \"d\":", dValue, ok)
	cValue, ok := m2["c"]
	fmt.Println("key \"c\":", cValue, ok)

	m3 := make(map[string]bool, 100) // arrange memory for 100 elements
	fmt.Println(m3)
	// порівняння m1 == m2 - заборонено
	//m1["d"] = 5 // panic: assignment to entry in nil map

	fmt.Println(m1["d"])
	fmt.Printf("len m2: %v\n", len(m2))
	fmt.Println(m2)
	delete(m2, "a")
	fmt.Println(m2)

	// збереження лише унікальних значень
	set := map[int]bool{
		1: true,
		2: true,
		3: true,
		// 2: false, - ключ не унікальний. код не компілюється

	}

	m5 := map[string]map[string]int{
		"aa": {
			"f": 35,
			"c": 41,
		},
		"bb": {
			"f": 67,
		},
	}
	fmt.Println(set)
	fmt.Println(m5)

}

/*
недозволені ключи: slice, map
map[[]string]string - заборонено
map[map[string]string]int - заборонено
*/
