package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Elf struct {
	Calories    []int
	CaloriesSum int
}

func main() {
	inputRaw, err := os.ReadFile("./in.txt")
	checkErr(err)

	var elves []Elf

	currentElf := Elf{}
	for _, line := range strings.Split(string(inputRaw), "\n") {
		if line == "" {
			elves = append(elves, currentElf)
			currentElf = Elf{}
			continue
		}

		calories, err := strconv.Atoi(line)
		checkErr(err)

		currentElf.Calories = append(currentElf.Calories, calories)
		currentElf.CaloriesSum += calories
	}
	elves = append(elves, currentElf)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].CaloriesSum > elves[j].CaloriesSum
	})

	fmt.Printf("Max Elf carries %d\n", elves[0].CaloriesSum)
	fmt.Printf("Max3 Elf carries %d\n", elves[0].CaloriesSum+elves[1].CaloriesSum+elves[2].CaloriesSum)
}
