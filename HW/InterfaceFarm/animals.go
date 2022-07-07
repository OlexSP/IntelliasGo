package main

import "fmt"

const ( // the obvious usage did not occur to me
	dogEatsPerKG = 2
	cowEatsPerKG = 25
	catEatsPerKG = 7
)

type Dog struct {
	name   string
	weight int
}

func (d Dog) getFoodIntake() int {
	return dogEatsPerKG * d.weight
}

func (d Dog) String() string {
	return fmt.Sprintf("%s the dog weighs %d", d.name, d.weight)
}

type Cat struct {
	name   string
	weight int
}

func (c Cat) getFoodIntake() int {
	return catEatsPerKG * c.weight
}

func (c Cat) String() string {
	return fmt.Sprintf("%s the cat weighs %d", c.name, c.weight)
}

type Cow struct {
	name   string
	weight int
}

func (c Cow) getFoodIntake() int {
	return cowEatsPerKG * c.weight
}

func (c Cow) String() string {
	return fmt.Sprintf("%s the cow weighs %d", c.name, c.weight)
}

type foodIntakeGetter interface { // did it just for practice according to the lecture
	getFoodIntake() int
}

type animal interface {
	foodIntakeGetter
	fmt.Stringer
}
