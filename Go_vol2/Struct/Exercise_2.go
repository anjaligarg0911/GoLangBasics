package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 7,
			color: "red",
		},
		fourWheel: true,
	}
	fmt.Println(t1)
	fmt.Println(t1.doors)

	s1 := sedan{
		vehicle: vehicle{
			doors: 7,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(s1)
	fmt.Println(s1.doors)
}
