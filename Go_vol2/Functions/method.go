// func (r receiver) identifier(parameters) (return(s)) { code }

// Method is a function attached to a type
//  add receiver
package main

import "fmt"

type person struct {
	name string
	age  int
}

func (s person) intro() {
	fmt.Println(s.name, s.age)
}
func main() {
	p1 := person{
		name: "anjali",
		age:  19,
	}

	p1.intro()
}
