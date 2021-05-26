package main

import "fmt"

func main() {
	x := []string{"name1", "name2", "name3"}
	fmt.Println(x)

	y := []string{"obj1", "obj2"}
	fmt.Println(y)

	z := [][]string{x, y}
	fmt.Println(z)
	fmt.Println((len(z)))
	fmt.Println(z[0])
	fmt.Println(z[1])
}
