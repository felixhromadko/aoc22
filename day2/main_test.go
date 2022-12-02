package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// X=loose(0), Y=draw (3), Z=win (6)
// A=rock(1), B=paper(2), C=scissors(3)
func Test_P2(t *testing.T) {
	table := []struct {
		A      string
		B      string
		Output int
	}{
		{"A", "X", 3}, // play=scissors(3)+loose(0)
		{"A", "Y", 4}, // play=rock(1)+draw(3)
		{"A", "Z", 8}, // play=paper(2)+win(6)

		{"B", "X", 1}, // play=rock(1)+loose(0)
		{"B", "Y", 5}, // play=paper(2)+draw(3)
		{"B", "Z", 9}, // play=scissors(3)+win(6)

		{"C", "X", 2}, //play=paper(2)+loose(0)
		{"C", "Y", 6}, // play=scissors(3)+draw(3)
		{"C", "Z", 7}, // play=rock(1)+win(6)
	}

	for _, s := range table {
		res := P2([]byte(fmt.Sprintf("%s %s\n", s.A, s.B)))
		assert.Equalf(t, s.Output, res, "Input: %s %s fails!", s.A, s.B)
	}
}
