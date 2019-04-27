package main

import "fmt"
import "strconv"
import "sort"

func digitRoot(a int) int {

	num := a
	pop := 0
	total := 0

	for num != 0 {
		pop = num % 10
		total += pop
		num /= 10
	}

	return total
}

func main() {
	test := []int{100, 22, 4, 11, 31, 103}
	var digitArry []int

	digitRootMap := make(map[string]int)

	for i := 0; i < len(test); i++ {
		digitArry = append(digitArry, digitRoot(test[i]))
	}

	fmt.Println(digitArry)

	for i := 0; i < len(test); i++ {

		digitRootMap[strconv.Itoa(test[i])] = digitArry[i]
	}

	fmt.Println(digitRootMap)

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range digitRootMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	fmt.Println(ss)

	var finalArry []int

	for x := 0; x < len(ss); x++ {

		for i := 0; i < len(ss)-1; i++ {
			if ss[i].Value == ss[i+1].Value {
				key1, _ := strconv.Atoi(ss[i].Key)
				key2, _ := strconv.Atoi(ss[i+1].Key)
				if key1 > key2 {
					temp := ss[i]
					ss[i] = ss[i+1]
					ss[i+1] = temp
				}

			}
		}
	}

	fmt.Println(ss)

	for _, kv := range ss {
		key, _ := strconv.Atoi(kv.Key)
		finalArry = append(finalArry, key)

	}

	fmt.Println(finalArry)

}
