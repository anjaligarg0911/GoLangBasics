package main

import (
	"fmt"
	"strconv"
)

var p int = 20 // Global Variable

func main() {
	var i int = 10
	var s string = "Anjali Garg"
	var b bool = true
	fmt.Printf("%T => %v\n", i, i)
	fmt.Printf("%T => %q\n", s, s)
	fmt.Printf("%T => %v\n", b, b)

	var j int
	j = 3
	fmt.Printf("%T => %v\n", j, j)

	k := 2
	fmt.Printf("%T => %v\n", k, k)

	fmt.Printf("%T => %v\n", p, p)

	var q string
	q = strconv.Itoa(i)
	fmt.Printf("%T => %q\n", q, q)
}
