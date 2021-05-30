// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "anjali",
		last:  "garg",
		age:   12,
	}

	p1.speak()
}

func (p person) speak() {
	fmt.Println(p.first)
	fmt.Println(p.last)
	fmt.Println(p.age)
}
