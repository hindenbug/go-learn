package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	targetSum := 0
	fmt.Println(triplets(array, targetSum))
}

func triplets(a []int, target int) [][]int {
	sort.Ints(a)
	triplet := [][]int{}

	for i := 0; i < len(a)-2; i++ {
		next := i + 1
		prev := len(a) - 1
		for next < prev {
			if a[prev]+a[i]+a[next] == target {
				triplet = append(triplet, []int{a[i], a[next], a[prev]})
				next++
				prev--
			} else if a[prev]+a[i]+a[next] < target {
				next++
			} else {
				prev--
			}
		}
	}

	return triplet
}
