package main

import "fmt"

func main() {
	a := []int{2, 4, 4, 5, 3}

	// pass slice parameter to function
	fmt.Println(sum(a...))

	//pass 0 parameter to a function
	fmt.Println(sum1("anjali", 2))
	fmt.Println(sum1("anjali"))

}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}

func sum1(s string, a ...int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(sum)
	return sum
}
