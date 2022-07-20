package main

import (
	"fmt"
	"sync"
)

//   go run -race *.go (can be used with go build, go test)
func main() {
	//var mu sync.Mutex
	var wg sync.WaitGroup
	num := 15000
	counter := 0

	wg.Add(num) // adding control sum 1500  for wg.Wait() out of for cycle

	for i := 0; i < num; i++ {
		go func(n int) {
			defer wg.Done() // marks as done
			//mu.Lock()
			counter++
			//mu.Unlock()
		}(i)
	}
	wg.Wait() // waits for all goroutines to be done
	fmt.Println("counter: ", counter)
}
