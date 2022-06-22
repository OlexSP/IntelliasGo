package main

import "fmt"

func main() {

	var b int = 1
	//var c bool //= true

	/*if b == 0 {
		fmt.Printf("condition has met")
	} else if c {
		fmt.Printf("'else if' has done")
	} else {
		fmt.Printf("'else' was done")
	}*/

	switch b {
	case 1:
		fmt.Printf("we met case 1\n")
		fallthrough // we must do the next case
	case 2:
		fmt.Printf("we met case 2")
	default:
		fmt.Printf("cases hasn't been met")
	}

	// continue - next iteration
	// break - stop the circle

}
