package main

import (
	"fmt"
)

func main() {
	array := []int{5, 1, 4, 2}

	fmt.Println(arrayOfProducts(array))
}

func arrayOfProducts(a []int) []int {
	result := make([]int, len(a))
	leftProd := make([]int, len(a))
	rightProd := make([]int, len(a))
	result[0] = 1

	leftProd[0] = 1
	for i := 1; i < len(a); i++ {
		leftProd[i] = a[i-1] * leftProd[i-1]
	}

	rightProd[len(a)-1] = 1
	for i := len(a) - 2; i >= 0; i-- {
		fmt.Println(i, rightProd)

		rightProd[i] = a[i+1] * rightProd[i+1]
	}

	for i := range result {
		result[i] = leftProd[i] * rightProd[i]
	}

	return result
}
