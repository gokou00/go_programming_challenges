package main

import "fmt"

func main() {
	test := make(map[string][]string)

	test["john"] = append(test["john"], "Dog")
	test["john"] = append(test["john"], "Cat")

	fmt.Println(test["john"][0])

}
