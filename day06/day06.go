package day06

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

type XY struct {
	X, Y int
}

func Part1(input string) int {
	grid := make(map[XY]string)
	var maxX, maxY int
	{
		var y int
		for line := range strings.SplitSeq(input, "\n") {
			if line == "" {
				break
			}
			var x int
			for s := range strings.SplitSeq(line, " ") {
				if s == "" {
					continue
				}
				grid[XY{x, y}] = s
				x++
				maxX = x
			}
			y++
			maxY = y
		}
	}

	var totz int
	for x := range maxX {
		var nums []int
		for y := range maxY - 1 {
			nums = append(nums, toInt(grid[XY{x, y}]))
		}
		op := grid[XY{x, maxY - 1}]
		var result int
		switch op {
		case "+":
			result = 0
			for _, n := range nums {
				result += n
			}
		case "*":
			result = 1
			for _, n := range nums {
				result *= n
			}
		default:
			panic("bad op")
		}
		totz += result
	}

	return totz
}

func Part2(input string) int {
	var data [][]byte

	lines := strings.Split(input, "\n")
	data = make([][]byte, len(lines[0])+2)

	for _, line := range lines {
		if line == "" {
			break
		}
		for x := range line {
			data[x] = append(data[x], line[x])
		}
		data[len(line)] = append(data[len(line)], ' ')
	}

	//for i, line := range data {
	//	fmt.Println(i, ":", string(line))
	//}

	var totz int
	var nums []int
	var op byte
	for _, line := range data {
		if strings.TrimSpace(string(line)) == "" {
			if len(nums) == 0 {
				continue
			}
			switch op {
			case '+':
				val := 0
				for _, n := range nums {
					val += n
				}
				totz += val
			case '*':
				val := 1
				for _, n := range nums {
					val *= n
				}
				totz += val
			default:
				panic("bad op " + string(op) + " numbers " + fmt.Sprint(nums))
			}
			op = '~'
			nums = nums[:0]
			continue
		}
		if bytes.HasSuffix(line, []byte("+")) {
			op = '+'
			line = bytes.TrimSuffix(line, []byte("+"))
		}
		if bytes.HasSuffix(line, []byte("*")) {
			op = '*'
			line = bytes.TrimSuffix(line, []byte("*"))
		}
		nums = append(nums, toInt(strings.TrimSpace(string(line))))
		//fmt.Println(toInt(strings.TrimSpace(string(line))))
	}

	return totz
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	must(err)
	return i
}
