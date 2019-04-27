package main

import "fmt"

func main() {
	n := 3

	length := n * n
	count := 1

	//var spiralArray []int

	var test [][]int

	test = make([][]int, n)

	for i := range test {
		test[i] = make([]int, n)
	}

	var spiralMatrix [][]int
	spiralMatrix = make([][]int, n)

	for i := range spiralMatrix {
		spiralMatrix[i] = make([]int, n)
	}

	for i := 0; i < len(test); i++ {
		for j := 0; j < len(test); j++ {
			test[i][j] = count
			count++
		}

	}

	//fmt.Println(test)

	var testArray []int

	for i := 1; i <= length; i++ {
		testArray = append(testArray, i)
	}

	fmt.Println(testArray)
	count = 0
	i := 0
	k := 0
	l := 0
	m := n

	for k < m && l < n {
		for i = l; i < n; i++ {
			//spiralArray = append(spiralArray, testArray[count])
			spiralMatrix[k][i] = testArray[count]
			count++
		}
		k++

		for i = k; i < m; i++ {
			//spiralArray = append(spiralArray, testArray[count])
			spiralMatrix[i][n-1] = testArray[count]
			count++
		}
		n--
		if k < m {
			for i = n - 1; i >= l; i-- {
				//spiralArray = append(spiralArray, testArray[count])
				spiralMatrix[m-1][i] = testArray[count]
				count++
			}
			m--
		}

		if l < n {
			for i = m - 1; i >= k; i-- {
				//spiralArray = append(spiralArray, testArray[count])
				spiralMatrix[i][l] = testArray[count]
				count++
			}
			l++
		}

	}

	fmt.Println(spiralMatrix)

}
