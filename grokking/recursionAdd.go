package main

import "fmt"

func add(a []int) int {
	x := 0
	if len(a) == 1 {
		return a[0]
	} else {
		x = a[len(a)-1]     //get last item
		a = a[0 : len(a)-1] //remove last item from slice
		return x + add(a)   //recursive call
	}
}

func main() {
	test := []int{1, 2, 3, 5, 7, 10}
	fmt.Println(add(test))
}
