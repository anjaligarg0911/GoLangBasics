package main

import "fmt"

func main() {
	f1(2, 4, sum)
	f1(2, 4, multiply)

}

func sum(x, y int) {
	fmt.Println(x + y)
}

func multiply(x, y int) {
	fmt.Println(x * y)
}

func f1(x int, y int, f func(a, b int)) {
	f(x, y)
}
