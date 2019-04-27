package main

import (
	"fmt"
	"sort"
)

func main() {

	testStr := []string{"p1", "p2", "p3", "p4"}

	sort.Sort(sort.Reverse(sort.StringSlice(testStr)))

	fmt.Println(testStr)

}
