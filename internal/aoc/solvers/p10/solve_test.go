package p10

import (
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 7
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestParseLights(t *testing.T) {
	for _, testCase := range []struct {
		input string
		want  int
	}{
		{input: "[.##.]", want: 6},
		{input: "[...#.]", want: 8},
		{input: "[.###.#]", want: 46},
	} {
		t.Run(testCase.input, func(t *testing.T) {
			assert.Equal(t, parseLights(testCase.input), testCase.want)
		})
	}
}

func TestParseButton(t *testing.T) {
	for _, testCase := range []struct {
		input string
		want  int
	}{
		{input: "(3)", want: 8},
		{input: "(1,3)", want: 10},
	} {
		t.Run(testCase.input, func(t *testing.T) {
			assert.Equal(t, parseButton(testCase.input), testCase.want)
		})
	}
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 0
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}

func TestPushButton(t *testing.T) {
	state := 0
	button := 8

	state = pushButton(state, button)
	assert.Equal(t, state, 8)

	state = pushButton(state, button)
	assert.Equal(t, state, 0)
}
