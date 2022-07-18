package main

import "fmt"

func main() {
	for i := 0; i < 15; i += 1 {
		fmt.Println(i, ":", tiangleRowsCount(i))
	}

}

func tiangleRowsCount(amt int) int {
	rowAmt, rowN := 0, 0
	//if amt < 1 {
	//	return 0
	//}

	for {
		amt -= rowAmt
		if amt < rowAmt+rowN+1 {
			return rowN
		}
		rowN++
		rowAmt += rowN
	}
}
