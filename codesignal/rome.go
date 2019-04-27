package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 48
	ans := 0
	str := ""
	div := 0

	for n != 0 {

		ans = n / 1000

		if ans != 0 {
			str += strings.Repeat("M", ans)
			div = 1000 * ans
			n = n % div

		}

		ans = n / 900

		if ans != 0 {
			str += strings.Repeat("CM", ans)
			div = 900 * ans
			n = n % div

		}

		ans = n / 800

		if ans != 0 {
			str += strings.Repeat("DCCC", ans)
			div = 800 * ans
			n = n % div

		}

		ans = n / 700

		if ans != 0 {
			str += strings.Repeat("DCC", ans)
			div = 700 * ans
			n = n % div

		}

		ans = n / 600

		if ans != 0 {
			str += strings.Repeat("DC", ans)
			div = 600 * ans
			n = n % div

		}

		ans = n / 500

		if ans != 0 {
			str += strings.Repeat("D", ans)
			div = 500 * ans
			n = n % div

		}

		ans = n / 400

		if ans != 0 {
			str += strings.Repeat("CD", ans)
			div = 400 * ans
			n = n % div

		}

		ans = n / 300

		if ans != 0 {
			str += strings.Repeat("CCC", ans)
			div = 300 * ans
			n = n % div

		}

		ans = n / 200

		if ans != 0 {
			str += strings.Repeat("CC", ans)
			div = 200 * ans
			n = n % div

		}

		ans = n / 100

		if ans != 0 {
			str += strings.Repeat("C", ans)
			div = 100 * ans
			n = n % div

		}

		ans = n / 90

		if ans != 0 {
			str += strings.Repeat("XC", ans)
			div = 90 * ans
			n = n % div

		}

		ans = n / 80

		if ans != 0 {
			str += strings.Repeat("LXXX", ans)
			div = 80 * ans
			n = n % div

		}

		ans = n / 70

		if ans != 0 {
			str += strings.Repeat("LXX", ans)
			div = 70 * ans
			n = n % div

		}

		ans = n / 60

		if ans != 0 {
			str += strings.Repeat("LX", ans)
			div = 60 * ans
			n = n % div

		}

		ans = n / 50

		if ans != 0 {
			str += strings.Repeat("L", ans)
			div = 50 * ans
			n = n % div

		}

		ans = n / 40

		if ans != 0 {
			str += strings.Repeat("XL", ans)
			div = 40 * ans
			n = n % div

		}

		ans = n / 30

		if ans != 0 {
			str += strings.Repeat("XXX", ans)
			div = 30 * ans
			n = n % div

		}

		ans = n / 20

		if ans != 0 {
			str += strings.Repeat("XX", ans)
			div = 20 * ans
			n = n % div

		}

		ans = n / 10

		if ans != 0 {
			str += strings.Repeat("X", ans)
			div = 10 * ans
			n = n % div

		}

		ans = n / 9

		if ans != 0 {
			str += strings.Repeat("IX", ans)
			div = 9 * ans
			n = n % div

		}

		ans = n / 8

		if ans != 0 {
			str += strings.Repeat("VIII", ans)
			div = 8 * ans
			n = n % div

		}

		ans = n / 7

		if ans != 0 {
			str += strings.Repeat("VII", ans)
			div = 7 * ans
			n = n % div

		}

		ans = n / 6

		if ans != 0 {
			str += strings.Repeat("VI", ans)
			div = 6 * ans
			n = n % div

		}

		ans = n / 5

		if ans != 0 {
			str += strings.Repeat("V", ans)
			div = 5 * ans
			n = n % div

		}

		ans = n / 4

		if ans != 0 {
			str += strings.Repeat("IV", ans)
			div = 4 * ans
			n = n % div

		}

		ans = n / 3

		if ans != 0 {
			str += strings.Repeat("III", ans)
			div = 3 * ans
			n = n % div

		}

		ans = n / 2

		if ans != 0 {
			str += strings.Repeat("II", ans)
			div = 2 * ans
			n = n % div

		}

		ans = n / 1

		if ans != 0 {
			str += strings.Repeat("I", ans)
			div = 1 * ans
			n = n % div

		}

	}

	fmt.Println(str)

}
