package main

import "fmt"

func main() {
	table := []int{1, 2, 6, 24, 120, 720}
	n := 1

	for _, x := range table {
		if x >= n {
			fmt.Println(x)
			return
		}
	}

	fmt.Println(-1)

}
