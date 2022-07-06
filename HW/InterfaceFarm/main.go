package main

import "fmt"

func main() {
	farmAnimals := []animal{
		Cat{name: "Dambass", weight: 4},
		Cat{name: "Dana", weight: 8},
		Cow{name: "Murka", weight: 400},
		Cow{name: "Mashka", weight: -600},
		Cow{name: "Mira", weight: 450},
		Dog{name: "Pirat", weight: 15},
		Dog{name: "Patron", weight: 12},
	}

	farmNeeds := farmConsume(farmAnimals)

	fmt.Printf("The farm needs %v kg of food.\n", farmNeeds)

}
