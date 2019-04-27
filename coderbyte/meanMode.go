package main

import "fmt"

func main() {
	arr := []int{4, 4, 4, 6, 2}
	arrMap := make(map[int]int)

	mode := 0
	mean := 1
	total := 0
	temp := 0

	for _, x := range arr {
		total += x
	}

	mean = total / len(arr)

	for _, x := range arr {
		arrMap[x]++
	}

	for key, val := range arrMap {
		if val > temp && val >= 2 {
			temp = val
			mode = key
		}
	}

	//fmt.Println(mode)
	//fmt.Println(mean)

	if mode == mean {
		fmt.Println(1)
		return
	}

	fmt.Println(0)
}
