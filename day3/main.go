package main

import (
	"fmt"
	set "github.com/deckarep/golang-set/v2"
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Rucksack struct {
	ItemsLeft  set.Set[uint8]
	ItemsRight set.Set[uint8]
}

func main() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	var rucksacks []Rucksack

	for _, s := range strings.Split(string(inputRaw), "\n") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		if len(s)%2 != 0 {
			log.Fatalln("invalid line")
		}

		r := Rucksack{ItemsLeft: set.NewSet[uint8](), ItemsRight: set.NewSet[uint8]()}

		for i := 0; i < len(s)/2; i++ {
			r.ItemsLeft.Add(s[i])
		}

		for i := len(s) / 2; i < len(s); i++ {
			r.ItemsRight.Add(s[i])
		}

		rucksacks = append(rucksacks, r)
	}

	// find mis-placed items in each rucksack
	var misplaced []rune
	for _, r := range rucksacks {
		intersect := r.ItemsLeft.Intersect(r.ItemsRight)
		if intersect.Cardinality() != 1 {
			panic("zero or more than one duplicates")
		}
		misplaced = append(misplaced, rune(intersect.ToSlice()[0]))
	}

	var sumOfMisplaced int
	for _, item := range misplaced {
		sumOfMisplaced += RuneToPriority(item)
	}

	fmt.Printf("sumOfMisplaced: %d\n", sumOfMisplaced) //8123

	// find unique among three
	var badgeSum int
	for i := 0; i < len(rucksacks)/3; i++ {
		r1 := rucksacks[i*3]
		r2 := rucksacks[i*3+1]
		r3 := rucksacks[i*3+2]

		allItemsR1 := r1.ItemsLeft.Union(r1.ItemsRight)
		allItemsR2 := r2.ItemsLeft.Union(r2.ItemsRight)
		allItemsR3 := r3.ItemsLeft.Union(r3.ItemsRight)

		badge := allItemsR1.Intersect(allItemsR2).Intersect(allItemsR3)
		if badge.Cardinality() != 1 {
			panic("could not find unique badge")
		}

		badgeSum += RuneToPriority(rune(badge.ToSlice()[0]))
	}

	fmt.Printf("sumOfBadges: %d\n", badgeSum)

}

// a=1, z=26
// A=7, Z=52
func RuneToPriority(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r) - 96
	}

	if r >= 64 && r <= 90 {
		return int(r) - 38
	}

	panic("invalid rune")
}

//sumOfMisplaced: 8123
//sumOfBadges: 2620
