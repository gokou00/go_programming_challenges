package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "xoxo"

	numOfX := strings.Count(str, "x")
	numOfO := strings.Count(str, "o")

	fmt.Println(numOfO == numOfX)

}
