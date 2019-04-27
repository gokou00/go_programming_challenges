package main

import (
	"fmt"
	"strings"
)

func main() {
	noun := "dKrqO"

	//strBuild := ""

	noun = strings.ToLower(noun)

	strArr := strings.Split(noun, "")
	strArr[0] = strings.ToUpper(strArr[0])

	noun = strings.Join(strArr, "")
	fmt.Println(noun)

}
