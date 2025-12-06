package p6

import (
	"fmt"
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 4277556
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestSolveColumnPartOne(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedResult int
	}{
		{[]string{"123", "45", "6", "*"}, 33210},
		{[]string{"328", "64", "98", "+"}, 490},
		{[]string{"51", "387", "215", "*"}, 4243455},
		{[]string{"64", "23", "314", "+"}, 401},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.input), func(t *testing.T) {
			assert.Equal(t, solveColumnPartOne(testCase.input), testCase.expectedResult)
		})
	}
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 3263827
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}

func TestSolveColumnPartTwo(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedResult int
	}{
		{[]string{"123", " 45", "  6", "*"}, 8544},
		{[]string{"328", "64 ", "98 ", "+"}, 625},
		{[]string{" 51", "387", "215", "*"}, 3253600},
		{[]string{"64 ", "23 ", "314", "+"}, 1058},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.input), func(t *testing.T) {
			assert.Equal(t, solveColumnPartTwo(testCase.input), testCase.expectedResult)
		})
	}
}

func TestTransformColsToWeirdOctopusNumbers(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedResult []int
	}{
		{[]string{"123", " 45", "  6"}, []int{1, 24, 356}},
		{[]string{"328", "64 ", "98 "}, []int{369, 248, 8}},
		{[]string{" 51", "387", "215"}, []int{32, 581, 175}},
		{[]string{"64 ", "23 ", "314"}, []int{623, 431, 4}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.input), func(t *testing.T) {
			assert.DeepEqual(t, transformColToWeirdOctopusNumbers(testCase.input), testCase.expectedResult)
		})
	}
}

func TestParseInputPreservingSpaces(t *testing.T) {
	input := `123 328  51 22 64 
 45 64  387 5  23 
  6 98  215 2  314
*   +   *   +  +  `
	want := [][]string{
		{"123", "328", " 51", "22", "64 "},
		{" 45", "64 ", "387", "5 ", "23 "},
		{"  6", "98 ", "215", "2 ", "314"},
		{"*", "+", "*", "+", "+"},
	}
	assert.DeepEqual(t, parseInputPreservingSpaces(input), want)
}
