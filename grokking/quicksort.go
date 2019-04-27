package main

import "fmt"

// this does not work

func quicksort(a []int) []int {
	var less []int
	var greater []int
	var final []int

	if len(a) < 2 {
		return a
	} else {
		pivot := a[0]

		for _, x := range a {
			if x <= pivot {
				less = append(less, x)
			} else {
				greater = append(greater, x)
			}
		}

		final = append(quicksort(less), pivot)
		final = append(final, quicksort(greater)...)

		return final

	}
}

func main() {
	fmt.Println(quicksort([]int{3, 5, 20, 15, 18}))
}
