package day01

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

func Part1(input string) int {
	zeroes := 0
	position := 50
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		// R --> 1, L --> -1
		var multiplier int
		switch line[0] {
		case 'L':
			multiplier = -1
		case 'R':
			multiplier = 1
		default:
			log.Panicln("invalid direction:", line[0])
		}
		clicks, err := strconv.Atoi(line[1:])
		must(err)
		position += clicks * multiplier
		position %= 100
		if position < 0 {
			position += 100
		}
		if position == 0 {
			zeroes++
		}
	}
	return zeroes
}

func Part2(input string) int {
	zeroes := 0
	position := 50
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		// R --> 1, L --> -1
		var multiplier int
		switch line[0] {
		case 'L':
			multiplier = -1
		case 'R':
			multiplier = 1
		default:
			log.Panicln("invalid direction:", line[0])
		}
		clicks, err := strconv.Atoi(line[1:])
		must(err)

		for range clicks {
			position += multiplier
			position %= 100
			if position < 0 {
				position += 100
			}
			if position == 0 {
				zeroes++
			}
		}
	}
	return zeroes
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
