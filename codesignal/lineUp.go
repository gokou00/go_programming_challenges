package main

import "fmt"

func main() {
	commands := "RLARA"

	total := 0
	var command []string

	for _, x := range commands {
		fmt.Println(string(x))

		if len(command) == 0 && string(x) == "L" || len(command) == 0 && string(x) == "R" {
			command = append(command, string(x))
			fmt.Println("total", total)
			fmt.Println("command", command)
			continue
		}

		if len(command) == 0 && string(x) == "A" {
			total++
			fmt.Println("total", total)
			fmt.Println("command", command)
			continue
		}

		if (string(x) == "L" || string(x) == "R") && (command[len(command)-1] == "L" || command[len(command)-1] == "R") {
			total++
			command[len(command)-1] = ""
			command = command[:len(command)-1]
			fmt.Println("total", total)
			fmt.Println("command", command)

		}

	}

	fmt.Println(total)

}
