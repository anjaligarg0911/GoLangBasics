package main

import "fmt"

func main() {
	s1 := f1()
	fmt.Println(s1)

	x := f2()
	fmt.Printf("%T\n", x)

	x1 := x()
	fmt.Println(x1)

	fmt.Println(x())

	fmt.Println(f2()())
}

func f1() string {
	s2 := "This is Anjali Garg"
	return s2

}

func f2() func() int {
	return func() int {
		return 40
	}
}
