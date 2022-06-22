package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 30; i++ {
		if i%5 == 0 {
			fmt.Println("i=", i)
		}
	}
	fmt.Println("===================")
	sum := 0
	for i := 30; i >= 0; i-- {
		if i%5 == 0 {
			fmt.Println("i=", i)
			sum += i
		}
	}
	fmt.Println("sum=", sum)

	/*counter := 0
	for {
		time.Sleep(time.Second) // pause for a second
		fmt.Println("I'm working")
		counter++
		if counter == 5 {
			break
		}

	}*/

	aSlice := []int{1, 4, 10, 15}
	for i := 0; i < len(aSlice); i++ {
		fmt.Println(i, aSlice[i])
		aSlice[i]++ // changing slice elements
	}
	fmt.Println(aSlice)

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for key, val := range m2 {
		fmt.Printf("key=%v, value =%v\n", key, val)
	}

}
