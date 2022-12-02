package main

import (
	"fmt"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inputRaw, err := os.ReadFile("./in.txt")
	checkErr(err)

	fmt.Println(inputRaw)
}
