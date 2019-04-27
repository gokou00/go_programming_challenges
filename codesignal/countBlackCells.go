package main

import "fmt"

func main() {
	//[m + n + gcd( m, n) - 2]
	n := 66666
	m := 88888

	ans := (m + n + gcd(m, n) - 2)

	fmt.Println(ans)

}

func gcd(a int, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a

}
