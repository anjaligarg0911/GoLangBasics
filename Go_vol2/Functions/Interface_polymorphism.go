// func (r receiver) identifier(parameters) (return(s)) { code }

// Method is a function attached to a type
//  add receiver
package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	std bool
}

type human interface {
	intro()
}

func (s person) intro() {
	fmt.Println("person", s.name, s.age)
}

func (s student) intro() {
	fmt.Println("student", s.name, s.age, s.std)
}

//  if any has this function intro() than it's interface type, and can call human1 function
func human1(h human) {
	switch h.(type) {
	case person:
		fmt.Println("person", h)
	case student:
		fmt.Println("student", h)
	default:
		fmt.Println("A Human")

	}
}
func main() {
	s1 := student{
		person: person{
			name: "anjali",
			age:  19,
		},
		std: true,
	}

	s1.intro()

	p1 := person{
		name: "anjali garg",
		age:  18,
	}

	p1.intro()

	human1(p1)
	human1(s1)

}
