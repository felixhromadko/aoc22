package main

import (
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const length = 14

func main() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	buffer := make([]string, length)
	// fill the buffer
	for i := 0; i < length; i++ {
		buffer[i] = "_"
	}

	for charNr, s := range strings.Split(string(inputRaw), "") {
		// shift + add to the buffer
		for i := 0; i < length-1; i++ {
			buffer[i] = buffer[i+1]
		}
		buffer[length-1] = s

		// if the buffer only contains unique strings, we found the start of the signal
		if charNr > length && len(lo.FindDuplicates[string](buffer)) == 0 {
			fmt.Printf("Found start of channel at %d\n", charNr+1)
			break
		}
	}
}
