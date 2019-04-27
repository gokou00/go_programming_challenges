package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(newFile)
	newFile.Close()
	fmt.Println("end")
}
