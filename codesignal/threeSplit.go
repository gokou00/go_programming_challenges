package main

import "fmt"

func main() {
	a := []int{0, -1, 0, -1, 0, -1}
	sum := 0

	for _, x := range a {
		sum += x
	}

	temp := sum % 3

	if temp != 0 {
		fmt.Println("0")
		return
	}

	sum /= 3
	//cut1 := 0
	//cut2 := 0
	count := 0
	sum1 := 0
	sum2 := 0

	for i := 0; i < len(a)-2; i++ {
		sum1 += a[i]

		if sum1 == sum {
			sum2 = 0

			//for cut2 in (cut1 + 1)..<a.count-1
			for j := i + 1; j < len(a)-1; j++ {
				sum2 += a[j]
				if sum2 == sum {
					count += 1
				}
			}

		}

	}
	fmt.Println(count)

}
