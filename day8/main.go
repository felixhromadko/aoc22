package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inputRaw, err := os.ReadFile("./inf.txt")
	checkErr(err)

	var trees [][]int

	for _, line := range strings.Split(string(inputRaw), "\n") {
		var treesInRow []int
		for _, heightS := range strings.Split(line, "") {
			height, err := strconv.Atoi(heightS)
			checkErr(err)

			treesInRow = append(treesInRow, height)
		}

		trees = append(trees, treesInRow)
	}

	// puzzle A
	var visibleTrees int
	for y := 0; y < len(trees); y++ {
	treeLoop:
		for x := 0; x < len(trees[y]); x++ {
			// if the tree is on the outer edge, it is visible
			if x == 0 || y == 0 || y == len(trees)-1 || x == len(trees[y])-1 {
				visibleTrees++
				continue
			}

			// scan up, down, left, right from each tree
			currentTreeHeight := trees[y][x]

			// scan left
			for scanX := x - 1; x >= 0; scanX-- {
				if currentTreeHeight <= trees[y][scanX] {
					break
				}

				if scanX == 0 {
					visibleTrees++
					continue treeLoop
				}
			}

			// scan right
			for scanX := x + 1; scanX < len(trees[y]); scanX++ {
				if currentTreeHeight <= trees[y][scanX] {
					break
				}

				if scanX == len(trees[y])-1 {
					visibleTrees++
					continue treeLoop
				}
			}

			// scan up
			for scanY := y - 1; scanY >= 0; scanY-- {
				if currentTreeHeight <= trees[scanY][x] {
					break
				}

				if scanY == 0 {
					visibleTrees++
					continue treeLoop
				}
			}

			// scan down
			for scanY := y + 1; scanY < len(trees); scanY++ {
				if currentTreeHeight <= trees[scanY][x] {
					break
				}

				if scanY == len(trees)-1 {
					visibleTrees++
					continue treeLoop
				}
			}
		}
	}

	println(visibleTrees)

	// puzzle B

	var maxTotalScore int
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			// if the tree is on the edge, the score will be 0. Ignore.
			if x == 0 || y == 0 || y == len(trees)-1 || x == len(trees[y])-1 {
				continue
			}

			currentTreeHeight := trees[y][x]

			// scan left
			var scoreLeft int
			for scanX := x - 1; x >= 0; scanX-- {
				scoreLeft++
				if currentTreeHeight <= trees[y][scanX] || scanX == 0 {
					break
				}
			}

			// scan right
			var scoreRight int
			for scanX := x + 1; scanX < len(trees[y]); scanX++ {
				scoreRight++
				if currentTreeHeight <= trees[y][scanX] {
					break
				}
			}

			// scan up
			var scoreUp int
			for scanY := y - 1; scanY > 0; scanY-- {
				scoreUp++
				if currentTreeHeight <= trees[scanY][x] {
					break
				}
			}

			// scan down
			var scoreDown int
			for scanY := y + 1; scanY < len(trees); scanY++ {
				scoreDown++
				if currentTreeHeight <= trees[scanY][x] {
					break
				}
			}

			totalScore := scoreUp * scoreDown * scoreLeft * scoreRight
			if totalScore > maxTotalScore {
				maxTotalScore = totalScore
			}
		}
	}
	println(maxTotalScore)
}
