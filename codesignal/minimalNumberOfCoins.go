package main

import "fmt"

func main() {
	coins := []int{1, 2, 10}
	price := 28
	numCoins := 0

	for i := len(coins) - 1; i >= 0; i-- {
		if price/coins[i] != 0 {
			numCoins += price / coins[i]
			price %= coins[i]

		}

	}

	fmt.Println(numCoins)

}
