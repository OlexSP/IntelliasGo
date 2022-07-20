package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	var sumAllSlices int
	var mu sync.Mutex //race condition check fails without mutex
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
			sliceSum := sum(n[i])
			mu.Lock()
			sumAllSlices += sliceSum
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Printf("result: %d\n", sumAllSlices)

}

func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
