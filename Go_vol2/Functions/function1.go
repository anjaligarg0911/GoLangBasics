package main

import "fmt"

func main() {
	fun1()

	fun2("Anjali")
	s1 := fun3("Anjali")
	fmt.Println(s1)

	first_name, last_name := fun4("Anjali", "Garg")
	fmt.Println(first_name)
	fmt.Println(last_name)

}

func fun1() {
	fmt.Println("hello You are in function1")
}

func fun2(name string) {
	fmt.Println("hello ", name)
}

func fun3(name string) string {
	return "Hello " + name
}

func fun4(name1 string, name2 string) (string, string) {
	first_name := name1
	last_name := name2

	return first_name, last_name
}
