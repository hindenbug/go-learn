package main

import (
	"fmt"
)

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}

	fmt.Println(moveElementToEnd(array, 2))
}

func moveElementToEnd(a []int, toMove int) []int {
	i := 0
	j := len(a) - 1

	for i < j {
		if a[j] == toMove {
			j--
			continue
		}

		if a[i] == toMove {
			a[i], a[j] = a[j], a[i]
		}
		i++
	}

	return a
}
