package main

import "fmt"

func main() {
	arr := []int{2, 4, 16, 24}

	isGeo := true
	isArith := true

	arith := arr[1] - arr[0]

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != arith {
			isArith = false
			break
		}
	}

	geo := arr[1] / arr[0]

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]/arr[i] != geo {
			isGeo = false
			break
		}
	}

	if isGeo {
		fmt.Println("Geometric")
		return
	} else if isArith {
		fmt.Println("Arithmetic")
		return
	} else {
		fmt.Println(-1)
		return
	}

}
