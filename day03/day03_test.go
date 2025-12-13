package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 357, Part1(input0))
	assert.Equal(t, 16854, Part1(input1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 3121910778619, Part2(input0))
	assert.Equal(t, 167526011932478, Part2(input1))
}
