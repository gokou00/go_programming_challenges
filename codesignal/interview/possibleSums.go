package main

import "fmt"

func main() {

	coins := []int{1, 2}
	quantity := []int{50000, 2}

	var arrayTest []int

	uniqueSums := make(map[int]int)

	/*

		combineCQ := make(map[int]int)

		for i := 0; i < len(coins); i++ {
			combineCQ[coins[i]] = quantity[i]
		}
	*/

	// build array

	for i := 0; i < len(coins); i++ {
		for j := 0; j < quantity[i]; j++ {
			arrayTest = append(arrayTest, coins[i])
		}

	}

	//fmt.Println(arrayTest)

	sum := 0

	for i := 0; i < len(arrayTest); i++ {
		sum += arrayTest[i]
	}

	//var dp [len(arrayTest)][sum + 1]bool

	dp := make([][]bool, len(arrayTest)+1)

	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	//fmt.Println(dp)

	for i := 0; i <= len(arrayTest); i++ {
		dp[i][0] = true
	}

	//	fmt.Println(dp)

	for i := 1; i <= len(arrayTest); i++ {
		dp[i][arrayTest[i-1]] = true
		for j := 1; j <= sum; j++ {
			if dp[i-1][j] == true {
				dp[i][j] = true
				dp[i][j+arrayTest[i-1]] = true
			}
		}
	}

	for j := 0; j <= sum; j++ {

		if dp[len(arrayTest)][j] == true {

			uniqueSums[j] = 0
		}
	}

	fmt.Println(len(uniqueSums) - 1)

}
