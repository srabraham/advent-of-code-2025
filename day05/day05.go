package day05

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

type LowHigh struct {
	Low, High int64
}

func intersect(a, b LowHigh) bool {
	if a.Low > b.High {
		return false
	}
	if a.High < b.Low {
		return false
	}
	return true
}

func merge(a, b LowHigh) LowHigh {
	return LowHigh{min(a.Low, b.Low), max(a.High, b.High)}
}

func Part1(input string) int64 {
	lines := strings.Split(input, "\n")
	var lowHighs []LowHigh
	var n int
	// first read the ranges
	for n = 0; n < len(lines); n++ {
		if lines[n] == "" {
			// done reading the ranges
			break
		}
		spl := strings.Split(lines[n], "-")
		lowHighs = append(lowHighs, LowHigh{
			Low:  toInt64(spl[0]),
			High: toInt64(spl[1]),
		})
	}
	var freshCount int64
	// now read the IDs
line:
	for n = n + 1; n < len(lines); n++ {
		if lines[n] == "" {
			break
		}
		id := toInt64(lines[n])
		for _, lowHigh := range lowHighs {
			if lowHigh.Low <= id && id <= lowHigh.High {
				freshCount++
				continue line
			}
		}
	}
	return freshCount
}

func toInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	must(err)
	return i
}

func Part2(input string) int64 {
	lines := strings.Split(input, "\n")
	var lowHighs []LowHigh
	var n int
	// first read the ranges
	for n = 0; n < len(lines); n++ {
		if lines[n] == "" {
			// done reading the ranges
			break
		}
		spl := strings.Split(lines[n], "-")
		lowHighs = append(lowHighs, LowHigh{
			Low:  toInt64(spl[0]),
			High: toInt64(spl[1]),
		})
	}

	success := true
	for success {
		lowHighs, success = mergeSomePair(lowHighs)
	}
	// nothing left to merge

	fmt.Println("lowHighs:", lowHighs)
	var count int64
	for _, lowHigh := range lowHighs {
		count += lowHigh.High - lowHigh.Low + 1
	}
	return count
}

func mergeSomePair(lowHighs []LowHigh) ([]LowHigh, bool) {
	for i := 0; i < len(lowHighs); i++ {
		for j := i + 1; j < len(lowHighs); j++ {
			if intersect(lowHighs[i], lowHighs[j]) {
				lowHighs = append(lowHighs, merge(lowHighs[i], lowHighs[j]))
				lowHighs = slices.Delete(lowHighs, j, j+1)
				lowHighs = slices.Delete(lowHighs, i, i+1)
				return lowHighs, true
			}
		}
	}
	return lowHighs, false
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
