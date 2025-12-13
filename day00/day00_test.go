package day00

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 3, Part1(input0))
	assert.Equal(t, 1018, Part1(input1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 6, Part2(input0))
	assert.Equal(t, 5815, Part2(input1))
}
