package p5

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 3
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestParseFreshRanges(t *testing.T) {
	input := file.MustReadToString("test.txt")
	freshRangesInput := strings.Split(input, "\n\n")[0]

	got := parseRanges(logger.New(false), freshRangesInput)
	want := []freshRange{{3, 5}, {10, 20}}
	assert.DeepEqual(t, got, want)
}

func TestIsFresh(t *testing.T) {
	input := file.MustReadToString("test.txt")
	freshRanges := parseRanges(logger.New(false), strings.Split(input, "\n\n")[0])

	testCases := []struct {
		ingredientID   int
		expectedResult bool
	}{
		{1, false},
		{5, true},
		{8, false},
		{11, true},
		{17, true},
		{32, false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", testCase.ingredientID), func(t *testing.T) {
			assert.Equal(t, freshRanges.isFresh(testCase.ingredientID), testCase.expectedResult)
		})
	}
}

func TestPartTwo(t *testing.T) {
	t.Skip("to come back to")
	input := file.MustReadToString("test.txt")
	want := 14
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}
