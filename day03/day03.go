package day03

import (
	_ "embed"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

func Part1(input string) int {
	tot := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		tot += joltageRec(line, 2)
	}
	return tot
}

func joltage(s string) int {
	// find the largest digit in [0, len-1)
	var ind0 int
	var byte0 byte
	for i := 0; i < len(s)-1; i++ {
		if s[i] > byte0 {
			byte0 = s[i]
			ind0 = i
		}
	}
	var byte1 byte
	for i := ind0 + 1; i < len(s); i++ {
		if s[i] > byte1 {
			byte1 = s[i]
		}
	}
	j := int(byte0-'0')*10 + int(byte1-'0')
	return j
}

func joltageRec(s string, digitsNeeded int) int {
	if digitsNeeded == 0 {
		return 0
	}
	var ind0 int
	var byte0 byte
	for i := 0; i < len(s)-digitsNeeded+1; i++ {
		if s[i] > byte0 {
			byte0 = s[i]
			ind0 = i
		}
	}

	return int(byte0-'0')*tenToThe(digitsNeeded-1) + joltageRec(s[ind0+1:], digitsNeeded-1)
}

func Part2(input string) int {
	tot := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		tot += joltageRec(line, 12)
	}
	return tot
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func tenToThe(exp int) int {
	result := 1
	for range exp {
		result *= 10
	}
	return result
}
