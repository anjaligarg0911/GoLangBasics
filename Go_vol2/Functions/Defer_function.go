package main

import "fmt"

func main() {
	// defer funtion will get return the result in last (excute in last)
	defer fun1()
	defer fun2()

	fun1()
	fun2()
}

func fun1() {
	fmt.Println("This is fun1")
}
func fun2() {
	fmt.Println("This is fun2")
}
