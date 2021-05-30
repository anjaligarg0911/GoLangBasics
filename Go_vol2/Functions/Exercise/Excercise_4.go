// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type square struct {
	len float32
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float32 {
	return s.len * s.len
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		len: 2,
	}

	fmt.Println(s.area())

	c := circle{
		radius: 2,
	}

	fmt.Println(c.area())

	s1 := square{
		len: 2,
	}

	c1 := circle{
		radius: 2,
	}

	info(s1)
	info(c1)

}
