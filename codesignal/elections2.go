package main

import "fmt"
import "sort"

func main() {
	votes := []int{3, 1, 1, 3, 1}
	k := 2

	//votes := make([]int, len(votes))
	//var winners []int

	count := 0
	sort.Ints(votes)

	if k == 0 {
		//	sort.Sort(votes)
		if votes[len(votes)-1] == votes[len(votes)-2] {
			fmt.Println("0")
			return
		}

	}

	//	copy(votes, votes)

	for i := 0; i < len(votes); i++ {
		//temp := 0
		//	sort.Ints(votes)

		if votes[i]+k < votes[len(votes)-1] {
			continue
		} else if votes[i]+k > votes[len(votes)-1] {
			count++
			continue
		}

		if len(votes) > 2 && i == len(votes)-1 {
			if votes[len(votes)-1] > votes[len(votes)-2] {
				count++
			}

		}

	}
	fmt.Println(count)
	//	fmt.Println(winners)

}
