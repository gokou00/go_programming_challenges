package main

import "fmt"
import "sort"

func main() {

	var test []string // this array holds all the map elements

	//var test2 [][]string

	var testKeysSorted []string

	count := 0

	testMap := make(map[string][]string)

	var colLengths []int

	testMap["Sauce"] = []string{"Salad", "Pizza", "Quesadilla"}
	testMap["Cheese"] = []string{"Quesadilla", "Sandwich"}
	testMap["Tomato"] = []string{"Salad", "Pizza", "Sandwich"}
	testMap["Salad"] = []string{"Salad", "Sandwich"}

	for key, _ := range testMap {
		testKeysSorted = append(testKeysSorted, key)
	}

	sort.Strings(testKeysSorted)

	//need to sort by key value first

	for key, _ := range testMap {
		sort.Strings(testMap[key])
	}

	fmt.Println(testMap)

	for _, x := range testKeysSorted {
		colLengths = append(colLengths, len(testMap[x])+1)

	}

	fmt.Println(colLengths)

	test2 := make([][]string, len(testMap))

	for i := range test2 {
		test2[i] = make([]string, colLengths[i])
	}

	fmt.Println(test2)

	fmt.Println(testMap, "testMap")
	fmt.Println(testKeysSorted, "testKeysSorted")

	for _, x := range testKeysSorted {
		test = append(test, x)
		for _, y := range testMap[x] {
			test = append(test, y)
		}
	}

	fmt.Println(test)

	for i := 0; i < len(test2); i++ {
		for j := 0; j < len(test2[i]); j++ {

			test2[i][j] = test[count]
			count++
		}
	}

	fmt.Println(test2)

}
