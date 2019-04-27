package main

import "fmt"

func main() {
	score1 := 0
	score2 := 0

	if score1 == score2 {
		fmt.Println("false")
		return
	}

	if score1 > 7 || score2 > 7 {
		fmt.Println("false")
		return
	}

	if score1 == 6 && score2 < 5 {
		fmt.Println("true")
		return
	}

	if score2 == 6 && score1 < 5 {
		fmt.Println("true")
		return
	}

	if score1 == 7 && score2 == 5 {
		fmt.Println("true")
		return
	}

	if score2 == 7 && score1 == 5 {
		fmt.Println("true")
		return
	}

	if score1 >= 5 && score2 >= 5 {
		if score1 == 7 || score2 == 7 {
			fmt.Println("true")
			return
		}

	}

	fmt.Println("false")
	return

}
