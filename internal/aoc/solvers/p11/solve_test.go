package p11

import (
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 5
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test2.txt")
	want := 2
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}
