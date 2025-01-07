package main

import "fmt"

// Person represents a person with basic attributes
type Person struct {
	name   string
	age    int
	gender string
}

// Introduce is a method of Person that returns a self-introduction string
func (p Person) Introduce() string {
	return fmt.Sprintf("Hello! My name is %s, I am %d years old and I am %s.",
		p.name, p.age, p.gender)
}

func main() {
	// Create a new Person instance
	user := Person{
		name:   "John",
		age:    25,
		gender: "male",
	}

	// Call the Introduce method and print the result
	fmt.Println(user.Introduce())
}
