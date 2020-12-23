package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}

	fmt.Println(longestPeak(array))
}

func longestPeak(a []int) int {
	peak := 0

	if len(a) < 3 {
		return peak
	}

	i := 1

	for i < len(a)-1 {
		if a[i-1] < a[i] && a[i] > a[i+1] {
			j, k := i-1, i+2

			for j >= 0 && a[j] < a[j+1] {
				j--
			}

			for k < len(a) && a[k] < a[k-1] {
				k++
			}

			if k-j-1 > peak {
				peak = k - j - 1
			}
			i = k

		} else {
			i++
			continue
		}
	}

	return peak
}
