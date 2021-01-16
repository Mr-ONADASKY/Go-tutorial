package main

import (
	"fmt"
	"go-functions/simplemath"
)

func main() {
	answer, err := simplemath.Divide(6, 2)
	if err != nil {
		fmt.Printf("An error occured %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}

	numbers := []float64{12.2, 14, 16, 22, 4}
	total := simplemath.Sum(numbers...)
	fmt.Printf("total of sum: %f\n", total)
}