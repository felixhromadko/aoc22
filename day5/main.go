package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Move struct {
	Amount int
	From   int
	To     int
}

// Staple is a FirstInLastOut Staple
type Staple []uint8

func (s *Staple) Pop() uint8 {
	i := len(*s) - 1
	x := (*s)[i]
	*s = append((*s)[:i], (*s)[i+1:]...)

	return x
}

func (s *Staple) Push(u uint8) {
	*s = append(*s, u)
}

func (s *Staple) PushBottom(u uint8) {
	*s = append([]uint8{u}, *s...)
}

func (s *Staple) GetTopBox() uint8 {
	return (*s)[len(*s)-1]
}

// Print prints the staple from the top most, to the bottom element
func (s *Staple) Print() {
	for _, u := range *s {
		fmt.Printf("%s", string(u))
	}
	fmt.Println()
}

func main() {
	Step1()
	Step2()
}

func Step1() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	boxRegex := regexp.MustCompile("(\\s{3}|\\[(.)]) ?")
	moveRegex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	lines := strings.Split(string(inputRaw), "\n")

	var i int
	staples := map[int]*Staple{}

	// parse staples
	for i < len(lines) {
		boxRegexMatches := boxRegex.FindAllStringSubmatch(lines[i], -1)
		i++

		if len(boxRegexMatches) == 0 {
			break
		}

		for boxNr, match := range boxRegexMatches {
			// if the match is empty, there is no box in this slot.
			if len(strings.TrimSpace(match[2])) == 0 {
				continue
			}

			// if the match is not empty, there is something in this slot. prepend to array
			boxLabel := match[2][0]

			// get staple
			staple, ok := staples[boxNr+1]
			if !ok {
				staple = &Staple{}
				staples[boxNr+1] = staple
			}
			staple.PushBottom(boxLabel)
		}
	}

	//parse moves
	var moves []Move
	for i < len(lines) {
		moveRegexMatches := moveRegex.FindAllStringSubmatch(lines[i], -1)
		i++
		if len(moveRegexMatches) != 0 {
			moveAmountS := moveRegexMatches[0][1]
			moveFromS := moveRegexMatches[0][2]
			moveToS := moveRegexMatches[0][3]

			moveAmount, err := strconv.Atoi(moveAmountS)
			checkErr(err)
			moveFrom, err := strconv.Atoi(moveFromS)
			checkErr(err)
			moveTo, err := strconv.Atoi(moveToS)
			checkErr(err)

			moves = append(moves, Move{
				Amount: moveAmount,
				From:   moveFrom,
				To:     moveTo,
			})
		}
	}

	// execute moves
	for _, move := range moves {
		for cnt := 0; cnt < move.Amount; cnt++ {
			box := staples[move.From].Pop()
			staples[move.To].Push(box)
		}
	}

	var labels string
	for i := 1; i <= len(staples); i++ {
		box := staples[i]
		labels += string(box.GetTopBox())
	}

	fmt.Println(labels)
}

func Step2() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	boxRegex := regexp.MustCompile("(\\s{3}|\\[(.)]) ?")
	moveRegex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	lines := strings.Split(string(inputRaw), "\n")

	var i int
	staples := map[int]*Staple{}

	// parse staples
	for i < len(lines) {
		boxRegexMatches := boxRegex.FindAllStringSubmatch(lines[i], -1)
		i++

		if len(boxRegexMatches) == 0 {
			break
		}

		for boxNr, match := range boxRegexMatches {
			// if the match is empty, there is no box in this slot.
			if len(strings.TrimSpace(match[2])) == 0 {
				continue
			}

			// if the match is not empty, there is something in this slot. prepend to array
			boxLabel := match[2][0]

			// get staple
			staple, ok := staples[boxNr+1]
			if !ok {
				staple = &Staple{}
				staples[boxNr+1] = staple
			}
			staple.PushBottom(boxLabel)
		}
	}

	//parse moves
	var moves []Move
	for i < len(lines) {
		moveRegexMatches := moveRegex.FindAllStringSubmatch(lines[i], -1)
		i++
		if len(moveRegexMatches) != 0 {
			moveAmountS := moveRegexMatches[0][1]
			moveFromS := moveRegexMatches[0][2]
			moveToS := moveRegexMatches[0][3]

			moveAmount, err := strconv.Atoi(moveAmountS)
			checkErr(err)
			moveFrom, err := strconv.Atoi(moveFromS)
			checkErr(err)
			moveTo, err := strconv.Atoi(moveToS)
			checkErr(err)

			moves = append(moves, Move{
				Amount: moveAmount,
				From:   moveFrom,
				To:     moveTo,
			})
		}
	}

	// execute moves
	for _, move := range moves {
		var boxes []uint8
		for cnt := 0; cnt < move.Amount; cnt++ {
			box := staples[move.From].Pop()
			boxes = append(boxes, box)
		}

		for i := len(boxes) - 1; i >= 0; i-- {
			staples[move.To].Push(boxes[i])
		}
	}

	var labels string
	for i := 1; i <= len(staples); i++ {
		box := staples[i]
		labels += string(box.GetTopBox())
	}

	fmt.Println(labels)
}
