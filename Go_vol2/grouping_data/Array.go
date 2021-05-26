package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 4
	fmt.Println(x)
	fmt.Println(len(x))
	// Diffrence between slices and Array :-
	// Unlike arrays, slices can be resized using the built-in append function.
	//  Further, slices are reference types, meaning that they are cheap to assign
	//  and can be passed to other functions without having to create a new copy of its underlying array
}
