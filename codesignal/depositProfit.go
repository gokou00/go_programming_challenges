package main

import "fmt"

func main() {
	//Passed but there's an error in the hidden test that violates one of the rules
	deposit := 50
	rate := 25
	threshold := 100

	var ans float32

	ans = float32(deposit)

	count := 0
	rateDec := float32(rate) / 100

	fmt.Println(rateDec)

	for ans < float32(threshold) && count <= 6 {

		ans = (rateDec * float32(ans)) + float32(ans)
		fmt.Println(ans)
		count++

	}

	fmt.Println(count)

}
