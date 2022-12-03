package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type SetOfItems map[uint8]struct{}
type Rucksack struct {
	ItemsLeft  SetOfItems
	ItemsRight SetOfItems
	AllItems   SetOfItems
}

func (r *Rucksack) ToString() {
	for u := range r.AllItems {
		print(string(u))
	}
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

		r := Rucksack{ItemsLeft: SetOfItems{}, ItemsRight: SetOfItems{}, AllItems: SetOfItems{}}

		for i := 0; i < len(s)/2; i++ {
			r.ItemsLeft[s[i]] = struct{}{}
			r.AllItems[s[i]] = struct{}{}
		}

		for i := len(s) / 2; i < len(s); i++ {
			r.ItemsRight[s[i]] = struct{}{}
			r.AllItems[s[i]] = struct{}{}
		}

		rucksacks = append(rucksacks, r)
	}

	// find mis-placed items in each rucksack
	var misplaced []rune
	for _, rucksack := range rucksacks {
		for s := range rucksack.ItemsRight {
			_, ok := rucksack.ItemsLeft[s]
			if !ok {
				continue
			}
			misplaced = append(misplaced, rune(s))
		}
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

		r1.ToString()
		println()
		r2.ToString()
		println()
		r3.ToString()
		println()

		for s := range r1.AllItems {
			_, hasR2 := r2.AllItems[s]
			_, hasR3 := r3.AllItems[s]

			if !hasR2 || !hasR3 {
				continue
			}

			badgeSum += RuneToPriority(rune(s))
		}
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
