package main

import "fmt"

func main() {
	min1 := 1
	min2_10 := 2
	min11 := 1
	s := 6
	//answer 14
	total := 0

	count := 1

	for s > 0 {
		if count == 1 {
			total += min1
			s -= min1
			count++

		} else if count > 1 && count <= 10 {
			total += min2_10
			s -= min2_10
			count++
		} else if count > 10 {
			total += min11
			s -= min11
			count++
		}

		if s < 0 {
			count--
			break
		}

		fmt.Println(s)

		//fmt.Println(total)
		//i = s

		//count++
	}

	//fmt.Println(total)
	count -= 1

	fmt.Println(count)

}
