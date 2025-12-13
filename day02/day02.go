package day02

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

func Part1(input string) int {

	totalz := 0

	for _, line := range strings.Split(strings.TrimSpace(input), ",") {
		lowHigh := strings.Split(line, "-")
		if len(lowHigh) != 2 {
			panic("wanted two numbers")
		}
		low, high := mustInt(lowHigh[0]), mustInt(lowHigh[1])

		for i := low; i <= high; i++ {
			digitCount := countDigits(i)
			if digitCount%2 != 0 {
				continue
			}
			split := tenToThe(digitCount / 2)
			if (i / split) == i%split {
				totalz += i
			}
		}
	}
	return totalz
}

func Part2(input string) int {

	totalz := 0

	for _, line := range strings.Split(strings.TrimSpace(input), ",") {
		lowHigh := strings.Split(line, "-")
		if len(lowHigh) != 2 {
			panic("wanted two numbers")
		}
		low, high := mustInt(lowHigh[0]), mustInt(lowHigh[1])

	nextNum:
		for num := low; num <= high; num++ {
			numStr := strconv.Itoa(num)

		nextPeriod:
			for period := 1; period < len(numStr); period++ {
				if len(numStr)%period != 0 {
					continue nextPeriod
				}
				for i := 0; i < period; i++ {
					testDigit := numStr[i]
					for j := i + period; j < len(numStr); j += period {
						if testDigit != numStr[j] {
							continue nextPeriod
						}
					}
				}
				totalz += num
				//log.Println("valid", num, "period", period)
				continue nextNum
			}

			//digitCount := countDigits(num)
			//if digitCount%2 != 0 {
			//	continue
			//}
			//split := tenToThe(digitCount / 2)
			//if (num / split) == num%split {
			//	totalz += num
			//}
		}
	}
	return totalz
}

func tenToThe(exp int) int {
	result := 1
	for range exp {
		result *= 10
	}
	return result
}

func countDigits(num int) int {
	digits := 0
	for i := 0; ; i++ {
		if num == 0 {
			return digits
		}
		digits++
		num = num / 10
	}
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	must(err)
	return i
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
