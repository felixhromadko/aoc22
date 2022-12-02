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

type MoveType int

const (
	MoveTypeRock MoveType = iota
	MoveTypePaper
	MoveTypeScissors
)

type Outcome int

const (
	OutcomeWin Outcome = iota
	OutcomeDraw
	OutcomeLoose
)

func parseMove(input uint8) MoveType {
	switch input {
	case 'A', 'X':
		return MoveTypeRock
	case 'B', 'Y':
		return MoveTypePaper
	case 'C', 'Z':
		return MoveTypeScissors
	default:
		panic("illegal move")
	}
}

func parseOutcome(input uint8) Outcome {
	switch input {
	case 'X':
		return OutcomeLoose
	case 'Y':
		return OutcomeDraw
	case 'Z':
		return OutcomeWin
	default:
		panic("illegal outcome")
	}
}

type Round struct {
	Move1 MoveType
	Move2 MoveType
}

func (r Round) score() int {
	if (r.Move1+1)%3 == r.Move2 {
		return 6 + int(r.Move2) + 1
	} else if r.Move1 == r.Move2 {
		return 3 + int(r.Move2) + 1
	} else {
		return int(r.Move2) + 1
	}
}

func main() {
	inputRaw, err := os.ReadFile("./in.txt")
	checkErr(err)

	p1(inputRaw)
	P2(inputRaw)

}

func p1(inputRaw []byte) {
	var rounds []Round
	for _, s := range strings.Split(string(inputRaw), "\n") {
		move1 := parseMove(s[0])
		move2 := parseMove(s[2])

		rounds = append(rounds, Round{move1, move2})
	}

	var score int
	for _, round := range rounds {
		score += round.score()
	}

	fmt.Printf("Our score if 2nd is move: %d\n", score)
}

func P2(inputRaw []byte) int {
	var rounds []Round
	for _, s := range strings.Split(string(inputRaw), "\n") {
		if s == "" {
			continue
		}

		move1 := parseMove(s[0])
		outcome := parseOutcome(s[2])

		var move2 MoveType
		if outcome == OutcomeDraw {
			move2 = move1
		} else if outcome == OutcomeWin {
			move2 = (move1 + 1) % 3
		} else {
			move2 = (move1 + 3 - 1) % 3
		}

		rounds = append(rounds, Round{move1, move2})
	}

	var score int
	for _, round := range rounds {
		score += round.score()
	}

	fmt.Printf("Our score if 2nd is outcome: %d\n", score)
	return score
}
