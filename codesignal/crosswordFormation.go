package main

import "fmt"

func main() {
	words := []string{"anaesthetist",
		"thatch",
		"ethnics",
		"sabulous"}
	count := 0

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			for k := 0; k < len(words); k++ {
				for l := 0; l < len(words); l++ {
					if i != j && i != k && i != l && j != k && j != l && k != l {
						count += check(words[i], words[j], words[k], words[l])
					}
				}
			}
		}
	}

	fmt.Println(count)

}

func check(a string, b string, c string, d string) int {
	count := 0
	for a1 := 0; a1 < len(a); a1++ {
		for a2 := a1 + 2; a2 < len(a); a2++ {
			for b1 := 0; b1 < len(b); b1++ {
				for b2 := b1 + 2; b2 < len(b); b2++ {
					for c1 := 0; c1 < len(c); c1++ {
						for d1 := 0; d1 < len(d); d1++ {
							c2 := c1 + (a2 - a1)
							d2 := d1 + (b2 - b1)
							if c2 < len(c) && d2 < len(d) {
								if a[a1] == b[b1] && a[a2] == d[d1] && c[c1] == b[b2] && c[c2] == d[d2] {
									count++
								}
							}
						}
					}
				}
			}
		}
	}

	return count

}
