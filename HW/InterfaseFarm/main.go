package main

import "fmt"

func main() {
	farmAnimals := []animal{
		Cat{name: "Dambass", weight: 4},
		Cat{name: "Dana", weight: 8},
		Cow{name: "Murka", weight: 400},
		Cow{name: "Mashka", weight: 600},
		Cow{name: "Mira", weight: 450},
		Dog{name: "Pirat", weight: 15},
		Dog{name: "Patron", weight: 12},
	}

	farmNeeds := farmConsume(farmAnimals)

	fmt.Printf("The farm needs %v kg of food.\n", farmNeeds)

}

func farmConsume(farm []animal) int {
	var sumFoodWeight, foodIntake int
	for _, v := range farm {
		foodIntake = v.getFoodIntake()
		if foodIntake <= 0 {
			fmt.Printf("%s has got improper weight.\n", v.String())
			return 0
		}
		sumFoodWeight += foodIntake
		fmt.Printf("%s, eats %d kg/month.\n", v.String(), foodIntake)
	}
	return sumFoodWeight
}
