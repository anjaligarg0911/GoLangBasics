package main

import "fmt"

func main() {
	s := struct {
		first_name string
		friends    map[string]int
		favDrinks  []string
	}{
		first_name: "anjali",
		friends: map[string]int{
			"anjali": 1,
			"chhavi": 2,
		},
		favDrinks: []string{
			"mazza", "shake",
		},
	}

	fmt.Println(s)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for i, v := range s.friends {
		fmt.Println(i, v)
	}
	for i, v := range s.favDrinks {
		fmt.Println(i, v)
	}
}
