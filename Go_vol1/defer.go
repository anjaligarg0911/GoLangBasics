package main

import "fmt"

func main() {

	defer print1("Hi Anjali")
	print2("Garg")

}

func print1(s string) {
	fmt.Println(s)
}

func print2(s string) {
	fmt.Println(s)
}
