package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	testVal := []int{1000, 1210, 10000, 500, 487}
	for _, v := range testVal {
		fmt.Println(v, " : ", F(v))
	}

}
func F(n int) int {
	const runningGoroutineAmt = 15 // sets number of running goroutine
	var maxNum, maxEvenAmt int
	primeSlice := generatePrimes(n)

	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, runningGoroutineAmt)
	lenPrimeSlice := len(primeSlice)

	for i := lenPrimeSlice - (lenPrimeSlice / 5); i < lenPrimeSlice; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int) {
			defer wg.Done()
			evenAmt := evenCounter(primeSlice[i])
			mu.Lock()
			if evenAmt >= maxEvenAmt && primeSlice[i] > maxNum {
				maxEvenAmt, maxNum = evenAmt, primeSlice[i]
			}
			mu.Unlock()
			<-sem
		}(i)
	}
	wg.Wait()

	return maxNum
}

func evenCounter(value int) int {
	counter := 0
	for value > 0 {
		if value%10%2 == 0 {
			counter++
		}
		value /= 10
	}
	return counter
}

func generatePrimes(value int) []int { //sieveOfEratosthenes
	var primeSlice []int
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			primeSlice = append(primeSlice, i)
		}
	}
	return primeSlice
}

//func IsPrimeSqrt(value int) bool {
//	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
//		if value%i == 0 {
//			return false
//		}
//	}
//	return value > 1
//}
//
//func closestPrime(value int) int {
//	res := 2
//	if value < 2 {
//		return res
//	}
//	for i := value; i >= 2; i-- {
//		if IsPrimeSqrt(i) {
//			res = i
//			break
//		}
//	}
//	return res
//}
