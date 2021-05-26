package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println("\n")
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Print(i, j)
		fmt.Println("\n")
	}

Loop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
			break Loop
		}
		fmt.Println("\n")
	}

	k := []int{1, 2, 3, 4, 5}
	for key, value := range k {
		fmt.Println(key, value)
	}
}
