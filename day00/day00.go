package day00

import (
	_ "embed"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

func Part1(input string) int {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
	}
	return -1
}

func Part2(input string) int {
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
	}
	return -1
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
