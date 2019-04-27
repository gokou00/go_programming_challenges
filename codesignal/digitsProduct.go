package main

import "fmt"
import "strconv"
import "sort"

func main() {
	product := 450
	num := ""

	var factorArray []int
	var factorArrayStr []int

	for i := 2; i < product; i++ {
		if product%i == 0 {
			factorArray = append(factorArray, i)
		}
	}

	fmt.Println(factorArray)

	for i := 0; i < len(factorArray); i++ {
		for j := 0; j < len(factorArray); j++ {
			if j == i {
				continue
			}

			if factorArray[i]*factorArray[j] == product {
				num += strconv.Itoa(factorArray[i])
				num += strconv.Itoa(factorArray[j])
				temp, _ := strconv.Atoi(num)
				factorArrayStr = append(factorArrayStr, temp)
				num = ""

			}

		}

	}

	sort.Ints(factorArrayStr)

	fmt.Println(factorArrayStr)

}
