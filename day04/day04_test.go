package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 13, Part1(input0))
	assert.Equal(t, 1409, Part1(input1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 43, Part2(input0))
	assert.Equal(t, 8366, Part2(input1))
}
