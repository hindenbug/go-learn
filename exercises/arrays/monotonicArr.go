package main

import (
	"fmt"
)

func main() {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}

	fmt.Println(isMonotonic(array))
}

func isMonotonic(a []int) bool {
	isMonotonicNonInc, isMonotonicNonDec := true, true

	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			isMonotonicNonDec = false
		}

		if a[i] < a[i-1] {
			isMonotonicNonInc = false
		}
	}
	return isMonotonicNonInc || isMonotonicNonDec
}
