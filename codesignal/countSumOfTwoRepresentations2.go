package main

import "fmt"

func main() {
	n := 93
	l := 24
	r := 58
	total := 0

	/*

		for i := l; i <= r; i++ {
			if i+i == n {
				total++
			}
			if i+(i+1) == n {
				total++
			}

		}
	*/

	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			if i+j == n {
				total++
			} else if i+j > n {
				break
			}

		}

	}

	fmt.Println(total)

}
