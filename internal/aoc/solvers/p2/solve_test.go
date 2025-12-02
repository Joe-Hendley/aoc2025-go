package p2

import (
	"fmt"
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 1227775554
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 4174379265
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}

func TestParseRange(t *testing.T) {
	input := "11-22"
	got := parseRange(input)
	assert.DeepEqual(t, got, idRange{11, 22})
}

func TestParseInput(t *testing.T) {
	input := "11-22,95-115"
	got := parseInput(input)
	assert.DeepEqual(t, got, []idRange{{11, 22}, {95, 115}})
}

func TestIsValidIDPartOne(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult bool
	}{
		{11, false}, {12, true}, {101, true}, {1001, true}, {1010, false}, {38593859, false},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", testCase.input), func(t *testing.T) {
			assert.Equal(t, isValidIDPartOne(testCase.input), testCase.expectedResult)
		})
	}
}

func TestFindInvalidIDsPartOne(t *testing.T) {
	testCases := []struct {
		input           idRange
		expectedResults []int
	}{
		{idRange{11, 22}, []int{11, 22}}, {idRange{95, 115}, []int{99}}, {idRange{998, 1012}, []int{1010}},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.input), func(t *testing.T) {
			assert.DeepEqual(t, testCase.input.findInvalidIDsPartOne(), testCase.expectedResults)
		})
	}
}

func TestIsValidIDPartTwo(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult bool
	}{
		{11, false}, {12, true}, {101, true}, {1001, true}, {1010, false}, {38593859, false}, {111, false}, {565656, false},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", testCase.input), func(t *testing.T) {
			assert.Equal(t, isValidIDPartTwo(testCase.input), testCase.expectedResult)
		})
	}
}

func TestFindInvalidIDsPartTwo(t *testing.T) {
	testCases := []struct {
		input           idRange
		expectedResults []int
	}{
		{idRange{11, 22}, []int{11, 22}},
		{idRange{95, 115}, []int{99, 111}},
		{idRange{998, 1012}, []int{999, 1010}},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase.input), func(t *testing.T) {
			assert.DeepEqual(t, testCase.input.findInvalidIDsPartTwo(), testCase.expectedResults)
		})
	}
}
