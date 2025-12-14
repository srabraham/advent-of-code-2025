package day05

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPart1(t *testing.T) {
	require.Equal(t, int64(3), Part1(input0))
	require.Equal(t, int64(617), Part1(input1))
}

func TestPart2(t *testing.T) {
	require.Equal(t, int64(14), Part2(input0))
	require.Equal(t, int64(338258295736104), Part2(input1))
}
