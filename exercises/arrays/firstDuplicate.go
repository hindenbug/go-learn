package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{2, 1, 5, 2, 3, 3, 4}

	fmt.Println(firstDuplicateValue(array))
	fmt.Println(firstDuplicateValue2(array))
}

func firstDuplicateValue(a []int) int {
	seen := make(map[int]int)
	result := -1

	for _, elem := range a {
		seen[elem] += 1

		if val, _ := seen[elem]; val > 1 {
			result = elem
			break
		}
	}

	return result
}

func firstDuplicateValue2(a []int) int {
	result := -1

	for _, elem := range a {
		val := int(math.Abs(float64(elem)))

		if a[val-1] < 0 {
			result = elem
			break
		}

		a[val-1] *= -1
	}

	return result
}
