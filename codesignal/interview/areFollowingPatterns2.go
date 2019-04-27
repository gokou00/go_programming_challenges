package main

import "fmt"

//import "sort"

func main() {

	strings := []string{"cat", "dog", "dog"}

	patterns := []string{"a", "b", "b"}

	stringsMap := make(map[string]string)
	patternsMap := make(map[string]string)

	if len(strings) != len(patterns) {
		fmt.Println("false")
		return
	}

	for i := 0; i < len(strings); i++ {
		stringsMap[strings[i]] = patterns[i]
	}

	for i := 0; i < len(strings); i++ {
		if stringsMap[strings[i]] != patterns[i] {
			fmt.Println("false")
			return
		}
	}

	for i := 0; i < len(strings); i++ {
		patternsMap[patterns[i]] = strings[i]
	}

	for i := 0; i < len(strings); i++ {
		if patternsMap[patterns[i]] != strings[i] {
			fmt.Println("false")
			return
		}
	}

	fmt.Println("true")
}
