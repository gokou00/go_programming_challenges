package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []string{"Cat", "cat"}

	sort.Strings(test)

	fmt.Println(test)

}
