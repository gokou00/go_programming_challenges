package main

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	//fmt.Println(x)
	//fmt.Println(temp)
	*x = *y
	*y = temp
}

func main() {
	x, y := 3, 4
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
