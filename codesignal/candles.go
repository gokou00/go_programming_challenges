package main

import "fmt"

func main() {
	candlesNumber := 15
	makeNew := 2
	total := candlesNumber

	leftovers := candlesNumber
	temp2 := 0

	if candlesNumber < makeNew {
		fmt.Println(candlesNumber)
		return
	}

	for leftovers/makeNew > 0 {
		//	fmt.Println(leftovers, "leftovers")

		temp := leftovers / makeNew
		temp2 = leftovers % makeNew

		total += temp
		//total += temp2
		//	fmt.Println(temp, "temp")
		//fmt.Println(temp2, "temp2")
		leftovers = temp + temp2

		//	leftovers += temp
		//leftovers /= makeNew
		//leftovers += temp2

	}

	fmt.Println(total)
	//fmt.Println(temp2)

}
