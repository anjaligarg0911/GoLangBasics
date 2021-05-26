package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var str string = "Anjali Garg"
	var str1 = []string{"A", "N", "J", "A", "L", "I"}

	fmt.Println("Type of s = ", reflect.TypeOf(str))
	fmt.Println("Length of s = ", len(str))
	fmt.Println("s = ", strings.ToUpper(str))
	fmt.Println("s = ", strings.ToLower(str))
	fmt.Println("s = ", strings.HasPrefix(str, "Anj"))
	fmt.Println("s = ", strings.HasSuffix(str, "rg"))
	fmt.Println("s = ", strings.Join(str1, ""))
	fmt.Println("s = ", strings.Repeat(str, 2))
	fmt.Println("s = ", strings.Contains(str, "a"))
	fmt.Println("s = ", strings.Index(str, "An"))
	fmt.Println("s = ", strings.Count(str, "a"))
	fmt.Println("s = ", strings.Replace(str, "a", "p", 6))
	fmt.Println("s = ", strings.Split(str, ""))
	fmt.Println("s = ", strings.Compare(str, strings.Join(str1, "")))
	fmt.Println("s = ", strings.TrimSpace(str))
	fmt.Println("s = ", strings.ContainsAny(str, "p & l"))

}
