package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 40, Part1(input0, 10))
	assert.Equal(t, 83520, Part1(input1, 1000))
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		_ = Part1(input1, 1000)
	}
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 25272, Part2(input0))
	assert.Equal(t, 1131823407, Part2(input1))
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		_ = Part2(input1)
	}
}
