package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 5, 6, 7, 7))
}

func sum(x ...int) int {
	sum := 0

	for i, v := range x {
		sum = sum + v

		fmt.Println("for postion", i, "sum is = ", sum)
	}
	return sum
}
