package main

import (
	"fmt"
	"sort"
	//	"strings"
)

type kv struct {
	Key   string
	Value int
}

func main() {
	// Might not need to do this create an array of each letter
	// create a map for char counts
	// create another map with the non repeating chars
	// create a function to get the index of where char is in string. Put as value in second map

	s := "ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof"

	if len(s) == 1 {
		fmt.Println(s)
		return
	}

	charCounts := make(map[string]int)
	nonRepeating := make(map[string]int)

	var ss []kv

	for _, x := range s {
		//	fmt.Println(string(x))
		charCounts[string(x)]++

	}

	isSingle := false

	for _, x := range charCounts {

		if x == 1 {
			isSingle = true
		}

	}

	if isSingle == false {
		fmt.Println("_")
		return
	}

	//fmt.Println(charCounts)

	for i, x := range charCounts {
		if x == 1 {
			nonRepeating[i] = findIndex(s, i)
		}
	}

	//fmt.Println(nonRepeating)

	for k, v := range nonRepeating {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	fmt.Println(ss[0].Key)

}

func findIndex(str string, letter string) int {

	for i, x := range str {
		if string(x) == letter {
			return i
		}
	}

	return -1

}
