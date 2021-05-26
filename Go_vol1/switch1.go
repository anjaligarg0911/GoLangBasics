package main

import "fmt"

func main() {
	var i interface{} = "anjali"

	switch i.(type) {
	case int:
		fmt.Println("Hi, This is Int")
	case float32:
		fmt.Println("Hi, This is Float32")
	case float64:
		fmt.Println("Hi, This is Float64")
	case string:
		fmt.Println("Hi, This is String")
	default:
		fmt.Println("OOPS... Sorry it's nothing")
	}

	var j string = "anjali"

	switch j {

	case "anjali":
		fmt.Println("Hi, This is String")
		fallthrough
	case "anjali1":
		fmt.Println("Hi, This is String")
	default:
		fmt.Println("OOPS... Sorry it's nothing")
	}
}
