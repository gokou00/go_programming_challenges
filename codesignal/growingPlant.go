package main

import "fmt"

func main() {
	upSpeed := 6
	downSpeed := 5
	desiredHeight := 10

	count := 0

	//oneDay := upSpeed - downSpeed
	total := upSpeed
	count++

	for total < desiredHeight {
		total -= downSpeed
		//fmt.Println(oneDay)

		total += upSpeed
		count++

	}

	fmt.Println(count)

}
