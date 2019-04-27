package custPkg

func Min(arry []float64) float64 {

	temp := 0.0

	temp = arry[0]

	for i := 1; i < len(arry); i++ {
		if temp > arry[i] {
			temp = arry[i]
		}
	}

	return temp

}

func Max(arry []float64) float64 {

	temp := 0.0

	temp = arry[0]

	for i := 1; i < len(arry); i++ {
		if temp < arry[i] {
			temp = arry[i]
		}
	}

	return temp

}
