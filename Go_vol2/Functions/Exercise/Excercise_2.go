// create a func with the identifier foo that
// 	takes in a variadic parameter of type int
// 	pass in a value of type []int into your func (unfurl the []int)
// 	returns the sum of all values of type int passed in
// create a func with the identifier bar that
// 	takes in a parameter of type []int
// 	returns the sum of all values of type int passed in

package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(foo(l...))

	l1 := []int{1, 2, 3, 4, 5, 5}
	fmt.Println(bar(l1))
}

func foo(x1 ...int) int {
	sum := 0
	for _, v := range x1 {
		sum = sum + v
	}
	return sum
}

func bar(x1 []int) int {
	sum := 0
	for _, v := range x1 {
		sum = sum + v
	}
	return sum

}
