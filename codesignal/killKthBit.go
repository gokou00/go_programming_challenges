package main

import (
	"fmt"
	//"strconv"
)

func main() {
	n := 37
	k := 3
	mask := ^(1 << uint(k-1))
	n &= mask

	fmt.Println(n)

}
