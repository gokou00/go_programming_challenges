package main

import "fmt"

func main() {
	lastNumberOfDays := 31
	//28 , 30 , 31
	var ans []int

	switch lastNumberOfDays {
	case 28:
		ans = append(ans, 31)
	case 30:
		ans = append(ans, 31)
	case 31:
		ans = append(ans, 28)
		ans = append(ans, 30)
		ans = append(ans, 31)

	}

	fmt.Println(ans)

}
