package day04

import (
	_ "embed"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

type XY struct {
	X, Y int
}

var directions = []XY{
	{-1, -1}, {0, -1}, {+1, -1},
	{-1, 0} /* start */, {+1, 0},
	{-1, +1}, {0, +1}, {+1, +1},
}

func Part1(input string) int {
	grid := make(map[XY]byte)

	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		for x := range line {
			grid[XY{x, y}] = line[x]
		}
	}

	tot := 0
	for xy, val := range grid {
		if val != '@' {
			continue
		}
		if removable(xy, grid) {
			tot++
		}
	}
	return tot
}

func removable(xy XY, grid map[XY]byte) bool {
	adjacentRolls := 0
	for _, dir := range directions {
		if grid[XY{xy.X + dir.X, xy.Y + dir.Y}] == '@' {
			adjacentRolls++
		}
	}
	return adjacentRolls < 4
}

func Part2(input string) int {
	grid := make(map[XY]byte)

	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		for x := range line {
			grid[XY{x, y}] = line[x]
		}
	}

	removedTotal := 0
	for {
		removedInRound := 0
		for xy, val := range grid {
			if val != '@' {
				continue
			}
			if removable(xy, grid) {
				grid[xy] = '.'
				removedInRound++
			}
		}
		if removedInRound == 0 {
			return removedTotal
		}
		removedTotal += removedInRound
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
