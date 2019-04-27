package main

import "fmt"

func main() {
	young := false
	beautiful := false
	loved := true

	if young && beautiful && !loved {
		fmt.Println("true")
		return
	}

	if loved && !beautiful {
		fmt.Println("true")
		return
	}

	if loved && !young {
		fmt.Println("true")
		return
	}

	if loved && !young && !beautiful {
		fmt.Println("true")
		return
	}

	fmt.Println("false")

}
