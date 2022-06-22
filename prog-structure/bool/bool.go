package main

import "fmt"

var (
	b1        bool = false
	b2        bool = true
	b3        bool
	milkPrice = 39.25
	beerPrice = 76.36
	myMoney   = 160
	carPrice  = 500000.00
)

func main() {
	fmt.Println(b3)
	fmt.Println(b1 == b2)
	fmt.Println(!b1)
	isPriceEqual := milkPrice == beerPrice
	fmt.Println(isPriceEqual)
	canBuyBoth := (milkPrice + beerPrice) < float64(myMoney) // int to float
	println(canBuyBoth)
	carMoreExpensive := carPrice < milkPrice || carPrice > beerPrice
	println(carMoreExpensive)
}
