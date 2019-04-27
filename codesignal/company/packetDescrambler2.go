//https://repl.it/@kacade/Spacebot-July-1-2017
package main

import (
	"fmt"
	//"sort"
	//	"strconv"
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
	//sortedArry := make([]int, len(seq))
	//copy(sortedArry, seq)

	//sort.Ints(sortedArry)
	//startIdx := sortedArry[0]
	//endIdx := sortedArry[len(sortedArry)-1]
	//fmt.Println(seq)
	//fmt.Println(sortedArry)

	//strSeq := strings.Join()

	/*
		for _, x := range seq {
			temp := strconv.Itoa(x)
			strSeq += temp
		}
	*/
	/*
		//check to see if seq doesn't have any gaps
		for i := startIdx; i <= endIdx; i++ {
			temp := strconv.Itoa(i)

			if strings.Contains(strSeq, temp) == false {
				fmt.Println("")
				return
			}

		}
	*/

	//orderMap := make(map[int][]string)
	orderMap := make(map[int]string)
	//	strBuilder := ""

	//numStr := make(map[string]int)

	teest := ""

	//building my map of indexs of seq to fragmentData
	for i := 0; i < len(seq); i++ {
		orderMap[seq[i]] += fragmentData[i] //append(orderMap[seq[i]], fragmentData[i])

		if float64(strings.Count(orderMap[seq[i]], fragmentData[i])) > half {
			teest += fragmentData[i]
		}

	}

	fmt.Println(orderMap)
	fmt.Println(teest)

	/*
		var orderedArry []int

		for key, _ := range orderMap {
			orderedArry = append(orderedArry, key)
		}
		sort.Ints(orderedArry)

		anotherMap := make(map[int]string)

		for _, x := range orderedArry {
			anotherMap[x] = strings.Join(orderMap[x], "")
		}
		fmt.Println(anotherMap, "anotherMap")

		for _, x := range orderedArry {
			strBuilder += anotherMap[x]
		}

		fmt.Println(strBuilder)

		//fmt.Println(strBuilder[0:4])
		mapcount := orderedArry[0]
		//testStr := ""
		//yetAnothermap := make(map[string]int)
		//	count := 0
		ending := len(anotherMap[mapcount])
		print(ending)
		reallyFinal := ""
		//	isfound := false
	*/
	/*

		for i, x := range strBuilder {

			//	fmt.Println(i, ending)
			count++

			//fmt.Println(strBuilder[i:ending])

			if float64(strings.Count(strBuilder[i:ending], string(x))) > half && strings.Contains(testStr, string(x)) == false {
				//yetAnothermap[string(x)] = 0
				testStr += string(x)
				isfound = true

			}

			if count == len(anotherMap[mapcount]) {
				count = 0
				mapcount++
				ending += len(anotherMap[mapcount])
				reallyFinal += testStr
				testStr = ""
				if isfound {
					isfound = false
				} else {
					fmt.Println("")
					return
				}

			}

		}
	*/

	//	fmt.Println(reallyFinal, "test")

	/*
		yetAnothermap := make(map[string]int)
		count := 0
		mapcount := startIdx


		for _,x := range strBuilder{
		    if count == len(anotherMap[mapcount]){

		    }

		    yetAnothermap[string(x)]++


		}
	*/

	/*

		for _, x := range strBuilder {
			numStr[string(x)]++

		}
	*/

	//test to see which chars have over 50% elements in fragment

	//finalStr := ""
	/*

		for i := 0; i < len(strBuilder); i++ {
			if float64(strings.Count(strBuilder, string(strBuilder[i]))) > half {
				finalStr += string(strBuilder[i])
				i = strings.LastIndex(strBuilder, string(strBuilder[i]))
			}

		}

	*/
	/*
		isenough := false
		for i := startIdx; i <= endIdx; i++ {
			//tempArry := orderMap[i]
			tempStr := anotherMap[i]
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
	*/

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
	//	if len(reallyFinal) == 0 {
	//		fmt.Println("")
	//		return
	//	}

	//fmt.Println(numStr)
	//fmt.Println(finalStr)

	//final check to see if # is the last char
	//	if string(reallyFinal[len(reallyFinal)-1]) == "#" {
	//		fmt.Println(reallyFinal)
	//	} else {
	//		fmt.Println("")
	//	}

}
