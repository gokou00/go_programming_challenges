package main

//testing 2d arrays

import "fmt"

func main() {

	//	var test [][]int

	count := 0

	a := [3][4]int{
		{1, 0, 0, 2},
		{0, 5, 0, 1},
		{0, 0, 3, 5},
	}

	rows := []int{}

	cols := []int{0, 2}

	//	romTemp := 0

	//	numCols := len(a[0]) - len(cols)

	var test [][]int

	if len(rows) != 0 {

		numRows := len(a) - len(rows)

		test = make([][]int, numRows) //rows

		//cols
		for i := range test {
			test[i] = make([]int, len(a[0]))
		}

		for i := 0; i < len(a); i++ {
			for j := 0; j < len(rows); j++ {
				if i == rows[j] {
					continue
				} else {
					fmt.Println(a[i])
					for d := 0; d < len(a[i]); d++ {
						test[count][d] = a[i][d]

					}
				}
				count++

			}

		}
	} else {
		fmt.Println("no rows")
		test = make([][]int, len(a)) //rows

		//cols
		for i := range test {
			test[i] = make([]int, len(a[0]))
		}

		//copy(test, a)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[0]); j++ {
				test[i][j] = a[i][j]
			}
		}

	}

	fmt.Println(test)

	//create new 2d array

	numCols := len(a[0]) - len(cols)

	test2 := make([][]int, len(test)) //rows

	//cols
	for i := range test2 {
		test2[i] = make([]int, numCols)
	}

	//skip := false

	count = 0

	for i := 0; i < len(test); i++ {
		for j := 0; j < len(test[0]); j++ {

			if contains(j, cols) {
				continue
			} else {
				test2[i][count] = test[i][j]
				count++
			}

		}
		count = 0

	}

	fmt.Println(test2)

}

func contains(x int, cols []int) bool {
	// found := false
	for i := 0; i < len(cols); i++ {
		if x == cols[i] {
			return true
		}

	}

	return false
}
