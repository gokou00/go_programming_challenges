package main

import (
	"fmt"
	"strconv"
)

func main() {
	//need a map [string]string to act as the key
	crypt := []string{"SEND", "MORE", "MONEY"}

	solution := [][]string{{"O", "0"},
		{"M", "1"},
		{"Y", "2"},
		{"E", "5"},
		{"N", "6"},
		{"D", "7"},
		{"R", "8"},
		{"S", "9"},
	}

	lookup := make(map[string]string)

	c1 := crypt[0]
	c2 := crypt[1]
	c3 := crypt[2]

	s1 := ""
	s2 := ""
	s3 := ""

	n1 := 0
	n2 := 0
	n3 := 0

	//create the map lookup

	//build the keys
	for i := 0; i < len(solution); i++ {
		//lookup[solution[i][0]]
		for j := 1; j < len(solution[0]); j++ {
			lookup[solution[i][0]] = solution[i][j]

		}
	}

	//	fmt.Println(lookup)

	for _, x := range c1 {
		s1 += lookup[string(x)]
	}

	for _, x := range c2 {
		s2 += lookup[string(x)]
	}

	for _, x := range c3 {
		s3 += lookup[string(x)]
	}

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	if len(s1) == 1 {
		if string(s1[0]) == "0" && string(s2[0]) == "0" && string(s3[0]) == "0" {
			fmt.Println("true")
			return
		}
	}

	//leading zero check
	if string(s1[0]) == "0" || string(s2[0]) == "0" || string(s3[0]) == "0" {
		fmt.Println("false")
		return

	}

	n1, _ = strconv.Atoi(s1)
	n2, _ = strconv.Atoi(s2)
	n3, _ = strconv.Atoi(s3)

	total := n1 + n2

	if total == n3 {
		fmt.Println("True")

		return
	}

	fmt.Println("false")
	return
}
