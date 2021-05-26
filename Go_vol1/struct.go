package main

import "fmt"

type city struct {
	money   int
	product int
	place   []string
}

type animal struct {
	Name   string
	origin string
}

type bird struct {
	animal
	speed  float32
	CanFly bool
}

func main() {
	var c city = city{
		money:   1234,
		product: 23,
		place: []string{
			"Rajasthan",
			"India",
		},
	}
	fmt.Println(c)
	fmt.Println(c.money)

	acity := struct{ name string }{name: "Rajasthan"}
	acity1 := &acity

	acity1.name = "GGN"

	fmt.Println(acity)
	fmt.Println(acity1)

	// /////////////////

	b := bird{}

	b.Name = "emu"
	b.origin = "Australia"
	b.speed = 46
	b.CanFly = true

	fmt.Println(b)

}
