package main

import "fmt"

func main() {
	// x := type{values}  composite literals
	x := []int{4, 5, 6, 8, 4}
	fmt.Println(x)

	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, v := range x {
		fmt.Println(i, v)
	}

	// slicing of a slice
	fmt.Println(x)
	fmt.Println(x[1:3])

	for i := 0; i < 3; i++ {
		fmt.Println(i, x[i])
	}

	// append the elements in a slice

	x = append(x, 67, 4, 5, 7)
	fmt.Println(x)

	// append a complete slice
	y := []int{2, 10, 5, 6, 78}
	x = append(y, x...)
	fmt.Println(x)

	// delete an element or multiple element from a slice
	x = []int{23, 45, 6, 8, 8, 9, 0}
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// just to check the function append
	// x = []int{23, 45, 6, 8, 8, 9, 0}
	// x = append(2, x[4:]...)
	// fmt.Println(x)

	// slice using make(only create slice, length, capacity)

	a := make([]int, 10, 11)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 878)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	//  create double size array/slice when the length > capacity (dynamically)
	a = append(a, 878)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
