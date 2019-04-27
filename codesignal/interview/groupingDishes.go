package main

import "fmt"
import "sort"

func main() {
	//dishes is a 2d array dishes [][]string
	//for each ingredient list the dish

	ingredMap := make(map[string][]string)
	ingredMap2 := make(map[string][]string)

	//dish := ""

	dishes := [][]string{{"Pasta", "Tomato Sauce", "Onions", "Garlic"},
		{"Chicken Curry", "Chicken", "Curry Sauce"},
		{"Fried Rice", "Rice", "Onions", "Nuts"},
		{"Salad", "Spinach", "Nuts"},
		{"Sandwich", "Cheese", "Bread"},
		{"Quesadilla", "Chicken", "Cheese"},
	}

	for _, x := range dishes {
		for i, y := range x {
			if i == 0 {
				continue
			} else {
				//fmt.Println(y)
				ingredMap[y] = []string{}
			}

		}
	}

	//fmt.Println(ingredMap)

	for x, _ := range ingredMap {
		ingredMap[x] = contains(x, dishes)
	}

	//fmt.Println(ingredMap)

	//need a way to remove dup Solved

	for key, value := range ingredMap {
		if len(value) >= 2 {
			ingredMap2[key] = value
		}
	}

	fmt.Println(ingredMap2)

	//need to sort by ingred and dish

	var test []string // this array holds all the map elements
	var testKeysSorted []string

	count := 0
	var colLengths []int

	for key, _ := range ingredMap2 {
		testKeysSorted = append(testKeysSorted, key)
	}

	sort.Strings(testKeysSorted)

	//need to sort by key value first

	for key, _ := range ingredMap2 {
		sort.Strings(ingredMap2[key])
	}

	fmt.Println(ingredMap2)

	for _, x := range testKeysSorted {
		colLengths = append(colLengths, len(ingredMap2[x])+1)

	}

	fmt.Println(colLengths)

	test2 := make([][]string, len(ingredMap2))

	for i := range test2 {
		test2[i] = make([]string, colLengths[i])
	}

	//	fmt.Println(test2)

	//	fmt.Println(testM, "testMap")
	//	fmt.Println(testKeysSorted, "testKeysSorted")

	for _, x := range testKeysSorted {
		test = append(test, x)
		for _, y := range ingredMap2[x] {
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

func contains(ingredient string, dishes [][]string) []string {
	var dishArray []string

	str := ""
	count := 0

	for _, x := range dishes {
		for i, y := range x {
			if i == 0 {
				str = y
			}

			if y == ingredient && i != 0 {
				dishArray = append(dishArray, str)
				count++
			}

		}
	}

	return dishArray

}
