package main

import "fmt"
import "sort"

func main() {
	votes := []int{3, 1, 1, 3, 1}
	k := 2

	votesCopy := make([]int, len(votes))
	var winners []int

	count := 0
	/*

		sort.Ints(votes)

		if votes[len(votes)-1] == votes[len(votes)-2] && k == 0 {

			fmt.Println("0")
			return

		}

	*/

	if k == 0 {
		sort.Ints(votes)
		if votes[len(votes)-1] == votes[len(votes)-2] {
			fmt.Println("0")
			return
		}

	}

	copy(votesCopy, votes)

	for i := 0; i < len(votesCopy); i++ {
		temp := 0
		sort.Ints(votesCopy)

		if votesCopy[i]+k < votesCopy[len(votesCopy)-1] {
			continue
		}

		temp = votesCopy[i]
		votesCopy[i] = votesCopy[i] + k

		sort.Ints(votesCopy)

		//	if temp < votesCopy[len(votesCopy)-1] {

		//		temp = votesCopy[len(votesCopy)-1]
		//	}

		//fmt.Println(votesCopy)

		if votesCopy[len(votesCopy)-1] > votesCopy[len(votesCopy)-2] {
			//fmt.Println(temp, votesCopy[len(votesCopy)-1])
			if winners == nil {
				count++
				winners = append(winners, temp)

				//	votesCopy = votesCopy[:len(votesCopy)-2]

			} else if greaterAll(votesCopy[len(votesCopy)-1], votesCopy[:len(votesCopy)-1]) && contains(votesCopy[len(votesCopy)-1], winners) == false {
				//	fmt.Println(temp, "temp")
				count++
				winners = append(winners, temp)
				//	votesCopy = votesCopy[:len(votesCopy)-2]
			}

		}

		copy(votesCopy, votes)
	}

	fmt.Println(count)
	fmt.Println(winners)

}

func contains(num int, arr []int) bool {
	isHere := false

	for _, x := range arr {
		if num == x {
			isHere = true
			return isHere
		}
	}

	return isHere
}

func greaterAll(num int, arr []int) bool {
	isGreater := true

	for _, x := range arr {
		if num <= x {
			isGreater = false
			return isGreater
		}
	}

	return true
}
