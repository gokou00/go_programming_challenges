package main

//create a map[int]int the key is the element in the array and the value is the index

import (
	"fmt"
	"sort"
)

type kv struct {
	Key   int
	Value int
}

func main() {

	hasDup := false
	test := []int{1}
	//	var dupArry []int

	var ss []kv

	testMap := make(map[int]int)

	testMap2 := make(map[int]int)

	for _, x := range test {
		testMap[x] = 0
	}

	for i, x := range test {

		if testMap[x] < 2 {
			testMap[x]++
		}

		if testMap[x] == 2 {
			testMap2[x] = i
			testMap[x] = 4
		}

	}

	fmt.Println(testMap)
	fmt.Println(test)
	fmt.Println(testMap2)

	fmt.Println(testMap2)

	for i, _ := range testMap {
		if testMap[i] > 1 {
			hasDup = true
			//	fmt.Println("hello")
			//	dupArry = append(dupArry, i)
		}

	}

	if !hasDup {
		fmt.Println("-1")
		return
	}

	for k, v := range testMap2 {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	fmt.Println(ss[0].Key)

}
