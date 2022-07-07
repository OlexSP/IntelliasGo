package main

import "fmt"

func farmConsume(farm []animal) int {
	var sumFoodWeight, foodIntake int
	for _, v := range farm {
		foodIntake = v.getFoodIntake()
		if foodIntake <= 0 {
			fmt.Printf("%s. It has got improper weight.\n", v.String())
			return 0
		}
		sumFoodWeight += foodIntake
		fmt.Printf("%s, eats %d kg/month.\n", v.String(), foodIntake)
	}
	return sumFoodWeight
}
