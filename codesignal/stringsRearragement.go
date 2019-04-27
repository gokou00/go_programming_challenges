package main

import (
	"fmt"
	"strings"
)

/*
This one is a bit of a doozy, in C++ I recommend two functions:

bool adjacentNodes(string1, string2) - compares two strings and only returns true if the two are different by one character. This function can terminate early if the number of different characters are greater than 2.

-bool stringsRearrangement (vector<string>):
-- do
-- set test to bool true
--- for every adjacent pair of words, test if adjacent
---- if NOT adjacent test false
-- while (next_permutation of (inputarray) and !test)
-- return test

I think this method is preferable to creating a graph because there are a lot of opportunities for early termination, so while there may be inputArray.size()! number of permutations, you'll run through them very quickly.


*/

func main() {

	inputArray := []string{"abc", "bef", "bcc", "bec", "bbc", "bdc"}

	inputArray3 := []string{"abc", "bef", "bcc", "bec", "bbc", "bdc"}

	if len(inputArray) == 1 {
		fmt.Println("false")
		return
	}

	countTest := 0

	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == inputArray3[i] {
			countTest++
		}
	}

	if countTest == len(inputArray) {
		fmt.Println("true")
		return
	}

	countequal := 0

	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i] == inputArray[i+1] {
			countequal++
		}
	}

	if countequal == len(inputArray)-1 {
		fmt.Println("false")
		return
	}

	inputArray2 := make([]string, len(inputArray))

	fmt.Println(len(inputArray2))

	copy(inputArray2, inputArray)

	//fmt.Println(inputArray2)

	//brute force permutations

	for m := 0; m < len(inputArray); m++ {

		for i := 0; i < len(inputArray); i++ {
			for j := 0; j < len(inputArray); j++ {
				//	if i == j {
				//		continue
				//	} else {
				inputArray2[i], inputArray2[j] = inputArray2[j], inputArray2[i]
				//	}

				fmt.Println(inputArray2)
				testg := holderf(inputArray2)

				if testg == true {
					fmt.Println("final true")
					return
				}
				//	copy(inputArray2, inputArray)
			}
			//copy(inputArray2, inputArray)
		}
	}

	fmt.Println("false")

}

func holderf(a []string) bool {

	count := 0
	for i := 0; i < len(a)-1; i++ {
		test := strTest(a[i], a[i+1])
		//fmt.Println(test)

		if test == true {
			count++
		}

	}
	if count == len(a)-1 {
		return true
	}

	return false
}

func strTest(a string, b string) bool {
	count := 0

	for _, x := range a {
		test := strings.ContainsAny(b, string(x))
		if test == false {
			count++
		}
	}

	if count > 1 {
		return false
	}

	count = 0

	for _, x := range b {
		test := strings.ContainsAny(a, string(x))
		if test == false {
			count++
		}
	}

	if count > 1 {
		return false
	}

	//return true
	count = 0

	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			count++
		}
	}

	//	fmt.Println(count)

	if count != 1 {
		return false
	}

	return true

}
