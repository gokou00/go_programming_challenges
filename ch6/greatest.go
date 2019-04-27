package main

import "fmt"

func great(args ...int) int {
	temp := 0

	for i, x := range args {
		if i == 0 || x < temp {
			temp = x
		}
	}

	return temp

}

func main() {
	ans := great(4, 5, 6, 3, 8, 1, 2)
	fmt.Println(ans)
}
