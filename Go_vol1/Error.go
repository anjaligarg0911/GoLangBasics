package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Negative value is not acceptable")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := sqrt(-5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result1, err1 := sqrt(64)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}
}
