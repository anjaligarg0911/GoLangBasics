package main

import "fmt"

func main() {

	var a = map[string]int{
		"SWM": 322201,
		"JPR": 302017,
		"GGN": 122001,
	}

	var m = map[[3]int]string{}
	var m1 = map[int]string{}

	a["Delhi"] = 100001 // add a key value to map
	fmt.Println(a)

	val, ok1 := a["Delhi"]
	fmt.Println(val, ok1)

	_, ok := a["Delhi"]
	fmt.Println(ok)

	fmt.Println(len(a))
	delete(a, "delhi")
	delete(a, "Delhi")
	// m2 = make(map[int]string)
	fmt.Println(a, m, m1)
}
