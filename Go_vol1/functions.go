package main

import "fmt"

type Employee struct {
	first_name string
	last_name  string
}

func (emp Employee) fullname() {
	fmt.Println(emp.first_name + " " + emp.last_name)
}

func fullname1() int {
	return 1
}

func allAll(args ...int) (int, int) {
	add := 0
	sub := 0

	for _, value := range args {
		add = add + value
		sub = sub - value
	}
	return add, sub
}

// Factorail program recursion
func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	var emp Employee = Employee{"Anjali", "Garg"}
	emp.fullname()
	emp2 := fullname1()
	fmt.Println(emp2)

	emp3, emp4 := allAll(10, 20, 30, 40, 50)
	fmt.Print(emp3, " ", emp4, " ")

	fmt.Println(fact(5))

	// closers

	a := 2

	squareNum := func() int {
		a = a * a
		return a
	}

	fmt.Println(squareNum())
	fmt.Println(squareNum())
}
