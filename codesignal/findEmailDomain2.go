package main

import (
	"fmt"
	"strings"
)

func main() {
	address := "\" \"@xample.org"

	idx := strings.LastIndex(address, "@")

	fmt.Println(address[idx+1:])
	

}
