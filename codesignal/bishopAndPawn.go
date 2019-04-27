package main

import "fmt"
import "strconv"

func main() {

	bishop := "h1"
	pawn := "a8"

	pos1 := 0
	pos2 := 0
	dist := 0
	bishopNum, _ := strconv.Atoi(string(bishop[1]))
	pawnNum, _ := strconv.Atoi(string(pawn[1]))

	if string(bishop[0]) == string(pawn[0]) {
		fmt.Println("False")
		return
	}

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for i, x := range alphabet {
		if x == string(bishop[0]) {
			pos1 = i
		}
	}

	for i, x := range alphabet {
		if x == string(pawn[0]) {
			pos2 = i
		}
	}

	posDist := 0

	if pos1 >= pos2 {
		fmt.Println(pos1 - pos2)
		posDist = pos1 - pos2
	} else {
		fmt.Println(pos2 - pos1)
		posDist = pos2 - pos1
	}

	//fmt.Println(pos1)
	//fmt.Println(pos2)

	if bishopNum <= pawnNum {
		dist = pawnNum - bishopNum
	} else {
		dist = bishopNum - pawnNum
	}

	fmt.Println(dist)

	if posDist == dist {
		fmt.Println("True")
		return
	}

	fmt.Println("false")
	return

}
