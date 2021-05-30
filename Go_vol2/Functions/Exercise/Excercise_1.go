// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results

package main

import "fmt"

func main() {
	a := foo()

	fmt.Println(a)

	a1, s := bar()

	fmt.Println(a1)
	fmt.Println(s)
}

func foo() int {
	return 8
}

func bar() (int, string) {
	return 13, "This is Anjali Garg"
}
