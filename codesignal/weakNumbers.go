//https://www.geeksforgeeks.org/count-divisors-n-on13/ if time is too long
//https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
package main

import (
	"fmt"
	"sort"
)

type kv struct {
	Key   int
	Value int
}

func main() {
	n := 4
	//var weaknessArr []int
	weaknessCount := 0

	weaknessMap := make(map[int]int)

	var finalArr []int

	divisorsMap := make(map[int]int)

	for i := 1; i <= n; i++ {
		divisorsMap[i] = weakness(i)
	}

	weaknessMap[1] = 0

	for i := 2; i <= n; i++ {
		for j := i - 1; j > 0; j-- {
			if divisorsMap[i] < divisorsMap[j] {
				weaknessCount++
			}

		}
		weaknessMap[i] = weaknessCount
		weaknessCount = 0
	}

	//fmt.Println(divisorsMap[403])
	//fmt.Println(weaknessMap[403])
	//fmt.Println(weaknessMap[])

	var ss []kv

	for k, v := range weaknessMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	largestWeakness := 0
	numOfWeakness := 0

	largestWeakness = ss[0].Value

	for _, v := range weaknessMap {
		if largestWeakness == v {
			numOfWeakness++
		}
	}

	finalArr = append(finalArr, largestWeakness)
	finalArr = append(finalArr, numOfWeakness)

	fmt.Println(finalArr)

}

func weakness(num int) int {
	count := 0

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}

	return count

}
