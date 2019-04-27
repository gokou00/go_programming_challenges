package main

import "fmt"

func main() {
	a := "abab"
	b := "abeb"

	if len(a) == 0 || len(b) == 0 {
		return
	}

	if len(a) != len(b) {
		fmt.Println("false")

	}

	if len(a) == 1 {
		return false
	}

	if len(a) == 2 {

		if string(a[0]) == string(a[1]) {
			//fmt.Println("false")
			return true
		} else if string(a[1])+string(a[0]) == b {
			return true
		} else {
			return false
		}
	}

	//create a map of both str

	mapA := make(map[string]int)
	mapB := make(map[string]int)

	for i := 0; i < len(a); i++ {
		mapA[string(a[i])] = 0
	}

	for i := 0; i < len(b); i++ {
		mapB[string(b[i])] = 0
	}

	var arrKey1 []string
	var arrKey2 []string

	for key, _ := range mapA {
		arrKey1 = append(arrKey1, key)

	}

	for key, _ := range mapB {
		arrKey2 = append(arrKey2, key)

	}

	fmt.Println(arrKey1)
	fmt.Println(arrKey2)

	//	fmt.Println(len(mapA))

	if len(arrKey1) == len(arrKey2) {

		for i := 0; i < len(arrKey1); i++ {
			bool1 := contains(arrKey2, arrKey1[i])
			if bool1 == false {
				fmt.Println("false")
			}
		}

	} else {
		fmt.Println("false")
	}

	if len(arrKey1) == len(arrKey2) {

		for i := 0; i < len(arrKey1); i++ {
			bool2 := contains(arrKey1, arrKey2[i])
			if bool2 == false {
				fmt.Println("false")
			}
		}

	} else {
		fmt.Println("false")
	}

	count := 0

	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			count++
		}
	}

	fmt.Println(count)

	if count <= 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

func contains(arr []string, str string) bool {
	for _, x := range arr {
		if x == str {
			return true
		}
	}

	return false
}
