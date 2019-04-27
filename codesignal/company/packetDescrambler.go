//https://repl.it/@kacade/Spacebot-July-1-2017
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	seq := []int{1, 1, 2, 2, 2, 4, 4}
	//seq := []int{4, 4, 0, 0, 0, 5, 5, 5}
	fragmentData := []string{"+",
		"+",
		"A",
		"A",
		"B",
		"#",
		"#"}
	n := 3
	half := 0.0
	half = float64(n) / 2.0
	sortedArry := make([]int, len(seq))
	copy(sortedArry, seq)

	sort.Ints(sortedArry)
	startIdx := sortedArry[0]
	endIdx := sortedArry[len(sortedArry)-1]
	fmt.Println(seq)
	fmt.Println(sortedArry)

	strSeq := ""

	for _, x := range seq {
		temp := strconv.Itoa(x)
		strSeq += temp
	}

	//check to see if seq doesn't have any gaps
	for i := startIdx; i <= endIdx; i++ {
		temp := strconv.Itoa(i)

		if strings.Contains(strSeq, temp) == false {
			fmt.Println("")
			return
		}

	}

	orderMap := make(map[int][]string)
	strBuilder := ""

	//numStr := make(map[string]int)

	//building my map of indexs of seq to fragmentData
	for i := 0; i < len(seq); i++ {
		orderMap[seq[i]] = append(orderMap[seq[i]], fragmentData[i])

	}

	fmt.Println(orderMap)

	for i := startIdx; i <= endIdx; i++ {
		for _, x := range orderMap[i] {
			strBuilder += x
		}
	}

	fmt.Println(strBuilder)
	/*

		for _, x := range strBuilder {
			numStr[string(x)]++

		}
	*/

	//test to see which chars have over 50% elements in fragment

	finalStr := ""
	/*

		for i := 0; i < len(strBuilder); i++ {
			if float64(strings.Count(strBuilder, string(strBuilder[i]))) > half {
				finalStr += string(strBuilder[i])
				i = strings.LastIndex(strBuilder, string(strBuilder[i]))
			}

		}

	*/
	isenough := false
	for i := startIdx; i <= endIdx; i++ {
		tempArry := orderMap[i]
		tempStr := strings.Join(tempArry, "")
		for i := 0; i < len(tempStr); i++ {
			if float64(strings.Count(tempStr, string(tempStr[i]))) > half {
				finalStr += string(tempStr[i])
				i = strings.LastIndex(tempStr, string(tempStr[i]))
				isenough = true
			}

		}
		if isenough {
			isenough = false
		} else {
			fmt.Println("")
			return
		}
	}

	fmt.Println(finalStr)

	/*

		isInStr := false

		for _, val := range orderMap {
			for _, x := range val {
				if strings.Contains(finalStr, x) {
					isInStr = true
					break
				}

			}
			if isInStr {
				isInStr = false
			} else {
				fmt.Println("")
				return
			}

		}
	*/

	//if finalStr doesn't have any chars then return an empty string
	if len(finalStr) == 0 {
		fmt.Println("")
		return
	}

	//fmt.Println(numStr)
	//fmt.Println(finalStr)

	//final check to see if # is the last char
	if string(finalStr[len(finalStr)-1]) == "#" {
		fmt.Println(finalStr)
	} else {
		fmt.Println("")
	}

}
