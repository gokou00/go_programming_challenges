package main

//create a countx and county that equal 3.
// Sum all the values but only add to slice if countx and county = 3
import "fmt"

func main() {

	var sumArry []int
	//	temp := 0
	total := 0
	k := 3

	test2dArry := [][]int{{36, 0, 18, 9, 9, 45, 27},
		{27, 0, 54, 9, 0, 63, 90},
		{81, 63, 72, 45, 18, 27, 0},
		{0, 0, 9, 81, 27, 18, 45},
		{45, 45, 27, 27, 90, 81, 72},
		{45, 18, 9, 0, 9, 18, 45},
		{27, 81, 36, 63, 63, 72, 81},
	}

	//	counterx, countery := 0, 0
	count := 0

	length := len(test2dArry) - 2
	width := len(test2dArry[0]) - 2

	fmt.Println(length)
	fmt.Println(width)

	//var result [length][width]int

	result := make([][]int, length)

	for i := range result {
		result[i] = make([]int, width)
	}

	fmt.Println(result)
	fmt.Println(test2dArry)

	for i := 0; i < len(test2dArry)-k+1; i++ {
		for j := 0; j < len(test2dArry[0])-k+1; j++ {
			total = 0
			for p := i; p < k+i; p++ {
				for q := j; q < k+j; q++ {
					total += test2dArry[p][q]
				}
			}
			total = total / 9
			result[i][j] = total
			sumArry = append(sumArry, total)
			count++

		}
	}

	fmt.Println(result)
	fmt.Println(sumArry)
	fmt.Println(count)

}
