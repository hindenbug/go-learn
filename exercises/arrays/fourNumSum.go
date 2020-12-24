package main

import (
	"fmt"
)

func main() {
	array := []int{7, 6, 4, -1, 1, 2}
	target := 16

	fmt.Println(fourNumSum(array, target))
}

func fourNumSum(a []int, target int) [][]int {
	pairSum := map[int][][]int{}
	quadruplets := [][]int{}

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			sum := a[i] + a[j]
			diff := target - sum
			if pairs, ok := pairSum[diff]; ok {
				for _, pair := range pairs {
					quad := append(pair, a[i], a[j])
					quadruplets = append(quadruplets, quad)
				}
			}
		}

		for k := 0; k < i; k++ {
			sum := a[i] + a[k]
			pairSum[sum] = append(pairSum[sum], []int{a[k], a[i]})
		}
	}

	return quadruplets
}
