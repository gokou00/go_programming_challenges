package main

import "fmt"

func main() {
	systemNames := []string{"stage_1", "stage_1", "stage_2", "dragon"}
	stepNumbers := []int{2, 1, 12, 111}

	testmap := make(map[string][]int)

	for i := 0; i < len(systemNames); i++ {
		testmap[systemNames[i]] = append(testmap[systemNames[i]], stepNumbers[i])

	}

	for _, x := range testmap {
		checker := isIncrease(x)
		if checker == false {
			fmt.Println("false")
			return
		}

	}

	fmt.Println("true")

	//fmt.Println(testmap)

}

func isIncrease(arr []int) bool {
	isIncreasing := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			continue
		} else {
			return false
		}
	}

	return isIncreasing

}
