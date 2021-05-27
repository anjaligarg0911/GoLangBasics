package main

import "fmt"

// struct is composite data type, allows us to composer together values of different type
type person struct {
	first_name string
	last_name  string
	icecream   []string
}

func main() {
	p1 := person{
		first_name: "Anjali",
		last_name:  "Garg",
		icecream:   []string{"Chocolate", "mango"},
	}

	p2 := person{
		first_name: "Roopal",
		last_name:  "Agrawal",
		icecream:   []string{"Chocolate", "mango"},
	}

	fmt.Println(p1)

	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}
	fmt.Println(p2)

	for i, v := range p2.icecream {
		fmt.Println(i, v)
	}

	//  map key => person.last_name , value => person
	map1 := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for k, v := range map1 {
		fmt.Println(k, v)
		for i, val := range v.icecream {
			fmt.Println(i, val)
		}
	}
}
