package main

import "fmt"

func main() {

	symbol := "5"

	switch symbol {
	case "0":
		fmt.Println("one")
	case "1":
		fmt.Println("two")
	case "2":
		fmt.Println("three")
	case "3":
		return "odd"
	case "4":
		return "even"
	case "5":
		return "odd"
	case "6":
		return "even"
	case "7":
		return "odd"
	case "8":
		return "even"
	case "9":
		return "odd"
	default:
		return "not a digit"
	}

}
