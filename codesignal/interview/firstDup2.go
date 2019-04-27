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
	test := []int{1, 1, 2, 2, 1}
	var dupArry []int

	var ss []kv

	testMap := make(map[int]int)

	testMap2 := make(map[int]int)

	for _, x := range test {
		testMap[x] = 0
	}

	for _, x := range test {
		testMap[x]++

	}

	fmt.Println(testMap)

	for i, _ := range testMap {
		if testMap[i] > 1 {
			hasDup = true
			//	fmt.Println("hello")
			dupArry = append(dupArry, i)
		}

	}

	if !hasDup {
		fmt.Println("-1")
		return
	}

	fmt.Println(dupArry)

	for _, x := range dupArry {
		testMap2[x] = 0
	}

	fmt.Println(testMap2)

	count := 0

	for x, _ := range testMap2 {

		for i, j := range test {

			//	fmt.Print(x, j)
			//	fmt.Println()

			if x == j {
				count++
				//fmt.Println(count, "count", x, j)
			}

			if count == 2 && x == j {
				testMap2[j] = i
				fmt.Println(j, i)
				fmt.Println(count)
				count = 0
				continue
			}

		}
		//	fmt.Println(count)

	}

	fmt.Println(testMap2)

	for k, v := range testMap2 {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	fmt.Println(ss[0].Key)

}
