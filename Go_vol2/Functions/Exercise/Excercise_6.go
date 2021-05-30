// pass a func into a func as an argument

package main

import "fmt"

func main() {
	g := func(x int) int {
		return 1
	}

	x := foo(g, 7)
	fmt.Println(x)
}

func foo(f func(x int) int, x int) int {
	x1 := f(x)

	x1++

	return x1

}
