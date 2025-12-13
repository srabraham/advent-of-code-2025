package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 1227775554, Part1(input0))
	assert.Equal(t, 30608905813, Part1(input1))
}

func BenchmarkPart1(b *testing.B) {
	for b.Loop() {
		assert.Equal(b, 30608905813, Part1(input1))
	}
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 4174379265, Part2(input0))
	assert.Equal(t, 31898925685, Part2(input1))
}

func BenchmarkPart2(b *testing.B) {
	for b.Loop() {
		assert.Equal(b, 31898925685, Part2(input1))
	}
}
