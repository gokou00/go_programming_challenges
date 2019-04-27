package main

import "fmt"

import "strconv"

func contains(arr []int, num int) bool {
	for i := 0; i < len(arr); i++ {
		if num == arr[i] {
			return true
		}
	}

	return false
}

func areSimilar(a []int, b []int) bool {

	//convert := ""

	mapA := make(map[string]int)
	mapB := make(map[string]int)

	for _, x := range a {
		s := strconv.Itoa(x)
		mapA[s] = 0

	}

	for _, x := range b {
		s2 := strconv.Itoa(x)
		mapB[s2] = 0

	}

	for _, x := range a {
		s := strconv.Itoa(x)
		mapA[s]++

	}

	for _, x := range b {
		s := strconv.Itoa(x)
		mapB[s]++

	}

	fmt.Println(mapA)
	fmt.Println(mapB)

	testA := true
	testB := true

	for i := 0; i < len(a); i++ {
		testA = contains(b, a[i])

	}

	fmt.Println(testA)

	if testA == false {
		return false
	}

	for i := 0; i < len(b); i++ {
		testB = contains(a, b[i])

	}
	fmt.Println(testB)

	if testB == false {
		return false
	}

	letter := ""

	for i := 0; i < len(a); i++ {
		letter = strconv.Itoa(a[i])
		if mapA[letter] != mapB[letter] {
			return false
		}
	}

	count := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	if count > 2 {
		return false
	} else {
		return true
	}

}

func main() {
	a := []int{3, 1, 2, 3, 2}
	b := []int{1, 1, 2, 3, 1}

	fmt.Println(areSimilar(a, b))

}
