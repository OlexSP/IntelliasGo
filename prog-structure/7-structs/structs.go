package main

import (
	"fmt"
	"strconv"
)

type User struct {
	id        string
	email     string
	firstName string
	lastName  string
	age       int
}

func generateUser(name string) User {
	return User{firstName: name, age: 35}
}

func printUserData(u User) {
	println("Name:", u.firstName)
	println("age", u.age)
}

func incrementUserAge(u *User) {
	u.age = u.age + 1
}

func (u User) FullName() string {
	return u.firstName + " " + u.lastName + " " + strconv.Itoa(u.age)
}

func (u *User) incrementAge() {
	u.age = u.age + 1
}

func main() {

	user1 := User{
		id:        "1",
		email:     "vasya@gmail.com",
		firstName: "Alex",
		lastName:  "Al",
		age:       40,
	}

	user2 := generateUser("Boris")

	fmt.Println(user1)
	fmt.Println(user2)

	printUserData(user1)
	incrementUserAge(&user1)
	printUserData(user1)

	fmt.Println(user1.FullName())
	user1.incrementAge()
	fmt.Println(user1.FullName())
}
