package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 21, Part1(input0))
	assert.Equal(t, 1622, Part1(input1))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 40, Part2(input0))
	assert.Equal(t, 10357305916520, Part2(input1))
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		_ = Part1(input1)
	}
}
