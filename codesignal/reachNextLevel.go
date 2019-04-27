package main

import "fmt"

func main() {
	experience := 16
	threshold := 23
	reward := 16

	total := experience + reward

	if total >= threshold {
		fmt.Println("true")
		return
	}

	fmt.Println("false")

}
