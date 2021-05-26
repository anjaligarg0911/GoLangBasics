package main

import "fmt"

func main() {

	//  map is basically key, value pair.

	// map[key]value
	x := map[string]string{
		"first_name": "abc",
		"last_name":  "xyz",
	}
	fmt.Println(x)
	fmt.Println(x["first_name"])
	fmt.Println(x["last_name"])

	v, ok := x["mid_name"]
	fmt.Println(v, ok)

	if v, ok := x["last_name"]; ok {
		fmt.Println("If mid_name exist : ", v)
	}

	//  add element and range

	x["mid_name"] = "pqr"

	for k, v := range x {
		fmt.Println(k, v)
	}

	// delete a key from a map
	fmt.Println(x)
	delete(x, "mid_name")
	fmt.Println(x)
}
