package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	for i := range n {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("slice %d: %d\n", i, sum(n[i]))
		}(i)
	}
	wg.Wait()

}

func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
