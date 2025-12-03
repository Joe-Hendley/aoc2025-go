package p3

import (
	"strings"
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 357
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestSolveLinePartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	lines := strings.Split(input, "\n")

	expectedResults := []int{
		98, 89, 78, 92,
	}

	for lineIndex, expectedResult := range expectedResults {
		t.Run(lines[lineIndex], func(t *testing.T) {
			assert.Equal(t, solveLinePartOne(lines[lineIndex]), expectedResult)
		})
	}
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 3121910778619
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}

func TestSolveLinePartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	lines := strings.Split(input, "\n")

	expectedResults := []int{
		987654321111, 811111111119, 434234234278, 888911112111,
	}

	for lineIndex, expectedResult := range expectedResults {
		t.Run(lines[lineIndex], func(t *testing.T) {
			assert.Equal(t, solveLinePartTwo(lines[lineIndex]), expectedResult)
		})
	}
}
