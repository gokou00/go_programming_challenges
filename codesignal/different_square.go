package main

import "fmt"
import "strconv"

func main() {

	matrix := [][]int{{2, 5, 3, 4, 3, 1, 3, 2},
		{4, 5, 4, 1, 2, 4, 1, 3},
		{1, 1, 2, 1, 4, 1, 1, 5},
		{1, 3, 4, 2, 3, 4, 2, 4},
		{1, 5, 5, 2, 1, 3, 1, 1},
		{1, 2, 3, 3, 5, 1, 2, 4},
		{3, 1, 4, 4, 4, 1, 5, 5},
		{5, 1, 3, 3, 1, 5, 3, 5},
		{5, 4, 4, 3, 5, 4, 4, 4}}

	//matrix := [][]int{{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 9, 9, 9, 2, 3, 9}}

	k := 2
	total := 0
	count := 0
	strMatrix := ""
	var strMatrixArray []string
	strMatrixMap := make(map[string]int)

	for i := 0; i < len(matrix)-k+1; i++ {
		for j := 0; j < len(matrix[0])-k+1; j++ {
			//total = 0
			for p := i; p < k+i; p++ {
				for q := j; q < k+j; q++ {
					total++
					//fmt.Print(matrix[p][q])
					strMatrix += strconv.Itoa(matrix[p][q])
				}
				//fmt.Println()
			}
			//total = total / 9
			//result[i][j] = total
			//sumArry = append(sumArry, total)
			strMatrixArray = append(strMatrixArray, strMatrix)
			strMatrixMap[strMatrix] = 0
			strMatrix = ""
			count++

		}
	}

	fmt.Println(count, "count")
	fmt.Println(total, "total")
	fmt.Println(strMatrixArray)
	fmt.Println(len(strMatrixMap))

}
