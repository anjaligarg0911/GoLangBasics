package main

import "fmt"

func main() {
	fun1()

	// Anonymous function
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println(x)
	}(40)

	//  function expression

	f := func() {
		fmt.Println("Hi")
	}

	f()

	f1 := func(x int) {
		fmt.Println(x)
	}
	f1(4)

}

func fun1() {
	fmt.Println("Not Anonymous function")
}
