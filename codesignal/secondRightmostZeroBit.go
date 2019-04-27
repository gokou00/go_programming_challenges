package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	n := 37
	fmt.Println(int(math.Pow((strings.LastIndex(strconv.FormatInt(int64(n), 2)[:index], "0") + 1), float64(2))))
	index := strings.LastIndex(strconv.FormatInt(int64(n), 2), "0")

}
