package p1

import (
	"fmt"
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 3
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 6
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}

func TestExtractCommand(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"L65", -65}, {"R34", 34},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("case %d input %s expect %d", i, testCase.input, testCase.expected), func(t *testing.T) {
			assert.Equal(t, extractCommand(testCase.input), testCase.expected)
		})
	}
}

func TestCountClicks(t *testing.T) {
	cases := []struct {
		pos           int
		diff          int
		expectedCount int
	}{
		{50, -68, 1},
		{82, -30, 0},
		{52, 48, 1},
		{0, -5, 0},
		{95, 60, 1},
		{55, -55, 1},
		{0, 1, 0},
		{99, -99, 1},
		{0, 14, 0},
		{14, -82, 1},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("case %d pos %d diff %d expect %d", i, testCase.pos, testCase.diff, testCase.expectedCount), func(t *testing.T) {
			assert.Equal(t, countClicks(testCase.pos, testCase.diff), testCase.expectedCount)
		})
	}
}

func TestGetBoundedPosition(t *testing.T) {
	cases := []struct {
		pos         int
		expectedPos int
	}{
		{0, 0},
		{100, 0},
		{50, 50},
		{-18, 82},
		{-118, 82},
		{118, 18},
		{-518, 82},
		{1118, 18},
		{-1, 99},
		{-100, 0},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("case %d pos %d expect %d", i, testCase.pos, testCase.expectedPos), func(t *testing.T) {
			assert.Equal(t, getBoundedPosition(testCase.pos), testCase.expectedPos)
		})
	}
}
