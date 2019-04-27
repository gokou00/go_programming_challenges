package main

import (
	"fmt"
	"gopkg.in/karalabe/cookiejar.v2/collections/deque"
	"reflect"
)

//using interfaces https://stackoverflow.com/questions/14025833/range-over-interface-which-stores-a-slice

func main() {
	d := deque.New()
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"kate", "john"}
	graph["kate"] = []string{"joan"}

	fmt.Println(graph)

	d.PushLeft(graph["you"])

	var testArry []string

	addStrArry := d.PopLeft()

	s := reflect.ValueOf(addStrArry)

	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.Index(i))

		str := s.Index(i)

		testArry = append(testArry, str.String())
	}

	//	fmt.Println(addStrArry)
	/*
		for _, x := range addStrArry {
			fmt.Println(x)
		}
	*/

	fmt.Println(testArry)

	/*
		for i := 0; i < 3; i++ {
			d.PushLeft(i)
		}

		//fmt.Println(d)
		fmt.Println(d.PopRight())
		fmt.Println(d.PopRight())

	*/

}
