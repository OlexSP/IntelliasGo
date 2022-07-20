package main

import (
	"fmt"
	"sync"
)

func worker(k int) {
	fmt.Printf("worker %d starting\n", k)
	fmt.Printf("worker %d done\n", k)

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1) // counter has to be executed before a goroutine start
		go func(n int) {
			defer wg.Done() // marks as done
			worker(n)
		}(i)
	}
	wg.Wait() // waits for all goroutines to be done
}
