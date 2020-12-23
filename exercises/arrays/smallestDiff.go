package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	array1 := []int{-1, 5, 10, 20, 28, 3}
	array2 := []int{26, 134, 135, 15, 17}

	fmt.Println(smallestDiff(array1, array2))
}

func smallestDiff(a1, a2 []int) []int {
	sort.Ints(a1)
	sort.Ints(a2)
	result := []int{}
	var idx1, idx2 int
	minDiff := math.MaxInt64

	for idx1 < len(a1) && idx2 < len(a2) {
		n1, n2 := a1[idx1], a2[idx2]
		diff := int(math.Abs(float64(n1) - float64(n2)))

		if diff < minDiff {
			minDiff = diff
			result = []int{n1, n2}
		}

		if n1 < n2 {
			idx1++
		} else if n2 < n1 {
			idx2++
		} else {
			result = []int{n1, n2}
			break
		}
	}

	return result
}
