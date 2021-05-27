package main

import "fmt"

// struct is composite data type, allows us to composer together values of different type
type person struct {
	first_name string
	last_name  string
	age        int
}

// Embedded struct
type student struct {
	person
	sdt bool
}

func main() {
	p1 := person{
		first_name: "Anjali",
		last_name:  "Garg",
		age:        18,
	}

	student1 := student{
		person: person{
			first_name: "Anjali",
			last_name:  "Garg",
			age:        18,
		},
		sdt: true,
	}

	// Anonymous struct
	// composite type :- type{value}
	p2 := struct {
		name string
		age  int
	}{
		name: "Anjali",
		age:  18,
	}

	fmt.Println(p2)
	fmt.Println(p1)
	fmt.Println(p1.first_name, p1.last_name)
	fmt.Println(student1)
	fmt.Println(student1.age, student1.first_name, student1.last_name, student1.sdt)
}
