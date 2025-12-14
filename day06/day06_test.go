package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 4277556, Part1(input0))
	assert.Equal(t, 4309240495780, Part1(input1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 3263827, Part2(input0))
	assert.Equal(t, 9170286552289, Part2(input1))
}
