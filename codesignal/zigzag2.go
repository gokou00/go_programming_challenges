package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestZig(A []int, size int) int {

	//n := len(A)
	//var Table [size][2]int

	//fmt.Println(len(A))

	Table := make([][]int, size)

	for i := range Table {
		Table[i] = make([]int, 2)
	}

	for i := 0; i < len(A); i++ {
		Table[i][0] = 1
		Table[i][1] = 1
	}

	fmt.Println(Table)

	result := 1

	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			if A[j] < A[i] && Table[i][0] < Table[j][1]+1 {
				Table[i][0] = Table[j][1] + 1
			}
			if A[j] > A[i] && Table[i][1] < Table[j][0]+1 {
				Table[i][1] = Table[j][0] + 1
			}
		}
		if result < max(Table[i][0], Table[i][1]) {
			result = max(Table[i][0], Table[i][1])
		}
		fmt.Println(result)
	}

	fmt.Println(Table)

	return result

}

func main() {
	test := []int{2, 1, 4, 4, 1, 4, 4, 1, 2, 0, 1, 0, 0, 3, 1, 3, 4, 1, 3, 4}
	result := longestZig(test, len(test))
	fmt.Println(result)
}
