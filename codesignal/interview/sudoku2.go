package main

//create a countx and county that equal 3.
// Sum all the values but only add to slice if countx and county = 3
import "fmt"

func main() {

	k := 3

	test2dArry := [][]string{{".", ".", ".", "1", "4", ".", ".", "2", "."},
		{".", ".", "6", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},
		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}

	dups := make(map[string]int)

	//check rows

	//https://stackoverflow.com/questions/28870784/passing-map-parameters-in-golang

	for i := 0; i < len(test2dArry); i++ {
		for j := 0; j < len(test2dArry[0]); j++ {
			dups[test2dArry[i][j]]++
		}
		isSudoku := verifySudoku(dups)
		if isSudoku == false {
			fmt.Println("Not Sudoku")
			return
		}
		fmt.Println("Before erased", dups)
		dups = resetTest(dups)
		fmt.Println("After erased", dups)
	}

	dups = resetTest(dups)

	//check cols
	for i := 0; i < len(test2dArry); i++ {
		for j := 0; j < len(test2dArry[0]); j++ {
			dups[test2dArry[j][i]]++
		}
		isSudoku := verifySudoku(dups)
		if isSudoku == false {
			fmt.Println("Not Sudoku")
			return
		}
		fmt.Println("Before erased", dups)
		dups = resetTest(dups)
		fmt.Println("After erased", dups)
	}

	dups = resetTest(dups)

	//checks the 3x3 sub matrices
	for i := 0; i < len(test2dArry)-k+1; i = i + 3 {
		for j := 0; j < len(test2dArry[0])-k+1; j = j + 3 {
			//total = 0
			for p := i; p < k+i; p++ {
				for q := j; q < k+j; q++ {
					dups[test2dArry[p][q]]++
				}
			}

			isSudoku := verifySudoku(dups)
			if isSudoku == false {
				fmt.Println("Not a sudoku")
				return
			}
			fmt.Println("Before erased", dups)
			dups = resetTest(dups)
			fmt.Println("After erased", dups)

		}
	}

	fmt.Println("Is a sudoku")

}

func verifySudoku(a map[string]int) bool {
	for i, x := range a {
		if x > 1 && i != "." {
			return false
		}
	}

	return true
}

func resetTest(a map[string]int) map[string]int {
	for i, _ := range a {
		a[i] = 0
	}

	return a

}
