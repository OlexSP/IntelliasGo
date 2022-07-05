package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

func (d Dog) getFoodIntake() int {
	return 10 / 5 * d.weight
}

func (d Dog) String() string {
	return fmt.Sprintf("%s the dog weighs %d", d.name, d.weight)
}

type Cat struct {
	name   string
	weight int
}

func (c Cat) getFoodIntake() int {
	return 7 * c.weight
}

func (c Cat) String() string {
	return fmt.Sprintf("%s the cat weighs %d", c.name, c.weight)
}

type Cow struct {
	name   string
	weight int
}

func (c Cow) getFoodIntake() int {
	return 25 * c.weight
}

func (c Cow) String() string {
	return fmt.Sprintf("%s the cow weighs %d", c.name, c.weight)
}

type foodIntakeGetter interface {
	getFoodIntake() int
}

type animal interface {
	foodIntakeGetter
	fmt.Stringer
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
