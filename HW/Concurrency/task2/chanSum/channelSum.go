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
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	ch := make(chan int, len(n))

	for i := range n {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sum(n[i], ch)
		}(i)
	}
	wg.Wait()
	close(ch)

	fmt.Printf("result: %d\n", chSum(ch))

}

func sum(slice []int, c chan int) {
	var sum int
	for _, v := range slice {
		sum += v
	}
	c <- sum
}

func chSum(c chan int) int {
	var chSum int
	for i := range c {
		chSum += i
	}
	return chSum
}
