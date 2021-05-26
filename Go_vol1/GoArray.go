package main

import "fmt"

func main() {
	var a = [5]int{3, 4, 5, 6, 7}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}

	var b = [3][3]int{{1, 2, 3}, {4, 5, 6}}
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			fmt.Print(b[j][k])
		}
		fmt.Print("\n")
	}

	//  Array slice

	var arr = [4]string{"abc", "def", "ghi", "jkl"}

	fmt.Println(arr[1:3])
	fmt.Println(len(arr[1:3]))
	fmt.Println(cap(arr[1:3]))

}
