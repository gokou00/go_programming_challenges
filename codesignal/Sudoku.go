package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 2, 5, 4, 6, 9, 8, 7},
		{4, 6, 5, 8, 7, 9, 3, 2, 1},
		{7, 9, 8, 2, 1, 3, 6, 5, 4},
		{9, 2, 1, 4, 3, 5, 8, 7, 6},
		{3, 5, 4, 7, 6, 8, 2, 1, 9},
		{6, 8, 7, 1, 9, 2, 5, 4, 3},
		{5, 4, 6, 9, 8, 1, 4, 3, 2},
		{2, 7, 3, 6, 5, 7, 1, 9, 8},
		{8, 1, 9, 3, 2, 4, 7, 6, 5},
	}

	//true

	testMap := make(map[int]int)
	k := 3
	count := 0

	for i := 1; i <= 9; i++ {
		testMap[i] = 0
	}
	//checks for each 3x3 sub matrix
	for i := 0; i < len(grid)-k+1; i = i + 3 {
		for j := 0; j < len(grid[0])-k+1; j = j + 3 {
			//total = 0
			for p := i; p < k+i; p++ {
				for q := j; q < k+j; q++ {
					//dups[test2dArry[p][q]]++
					testMap[grid[p][q]]++
					count++
				}
			}
			//fmt.Println(testMap)
			for _, val := range testMap {
				if val != 1 {
					fmt.Println("false 3x3")
					return
				}
			}

			for key, _ := range testMap {
				testMap[key] = 0
			}

		}
	}

	for key, _ := range testMap {
		testMap[key] = 0
	}

	//checks each row
	for _, x := range grid {
		for _, rows := range x {
			testMap[rows]++
		}

		for _, val := range testMap {
			if val != 1 {
				fmt.Println("false rows")
				return
			}
		}

		for key, _ := range testMap {
			testMap[key] = 0
		}

	}

	//check cols

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			testMap[grid[j][i]]++
		}

		for _, val := range testMap {
			if val != 1 {
				fmt.Println("false cols")
				return
			}
		}

		for key, _ := range testMap {
			testMap[key] = 0
		}

	}

	fmt.Println("true")
	fmt.Println(count)

}
