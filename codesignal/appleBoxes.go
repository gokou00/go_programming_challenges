package main

import "fmt"

func main() {
	k := 1
	red := 0
	yellow := 0

	for i := 1; i <= k; i++ {

		if i%2 == 0 {
			temp_red := i * i
			red += temp_red

		} else if i%2 == 1 {
			temp_yellow := i * i
			yellow += temp_yellow
		}

	}

	final := red - yellow

	fmt.Println(final)

}
