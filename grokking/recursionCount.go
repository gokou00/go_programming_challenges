package main

import "fmt"

func RecursiveCount(a []int) int {
	count := 0

	if len(a) == 0 {
		return count
	} else {
		count++
		fmt.Println("rec")
		return count + RecursiveCount(a[0:len(a)-1])
	}
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(RecursiveCount(test))

}
