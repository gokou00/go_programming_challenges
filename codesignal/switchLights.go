package main

import "fmt"

func main() {
	a := []int{1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1}

	for i := 0; i < len(a); i++ {

		if a[i] == 1 {
			a[i] = 0

			for j := i - 1; j >= 0; j-- {
				if a[j] == 0 {
					a[j] = 1
				} else {
					a[j] = 0
				}

			}

		}

	}

	fmt.Println(a)

}
