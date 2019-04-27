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

	Table := make([][]int, size) // rows

	//cols
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
		for j := i - 1; j >= 0; j-- {
			if A[j] < A[i] { //Math.max(Z[j][1]+1,Z[i][0])
				Table[i][0] = max(Table[j][1]+1, Table[i][0])
			}
			if A[j] > A[i] { //Math.max(Z[j][0]+1, Z[i][1])
				Table[i][1] = max(Table[j][0]+1, Table[i][1])
			}
		}
		//best = Math.max(best, Math.max(Z[i][0],Z[i][1]));
		result = max(result, max(Table[i][0], Table[i][1]))
		fmt.Println(result)
	}

	fmt.Println(Table)

	return result

}

func main() {
	test := []int{9, 8, 8, 5, 3, 5, 3, 2, 8, 6}
	result := longestZig(test, len(test))
	fmt.Println(result)
}
