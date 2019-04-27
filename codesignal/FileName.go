package main

import "fmt"
import "strconv"

func main() {
	names := []string{"a(1)",
		"a(6)",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a"}

	names_changed := []string{""}

	//finalNames = append(finalNames,names[0])
	temp := ""

	for i := 0; i < len(names); i++ {
		//finalNames = append(finalNames, checknames(names, names[i], i))
		names[i], temp = checknames(names, names[i], i, names_changed)
		names_changed = append(names_changed, temp)

	}

	fmt.Println(names)

}

func checknames(arr []string, name string, pos int, changed []string) (string, string) {
	//newStr := ""
	count := 0
	newName := ""
	for _, x := range changed {
		if x == name {
			count++
		}
	}

	for i := 0; i < pos; i++ {
		if name == arr[i] {
			count++
		}

	}

	if count == 0 {
		return name, ""
	}

	temp := strconv.Itoa(count)

	newName += name + "(" + temp + ")"

	for _, x := range arr {
		if x == newName {
			count++
			temp := strconv.Itoa(count)
			newName = name + "(" + temp + ")"
		}
	}

	return newName, name

}
