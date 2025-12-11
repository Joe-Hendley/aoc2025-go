package p8

import (
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/file"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
)

func TestPartOne(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 40
	solver := Solver{}
	solver.Init(true)

	got := solver.Part1(input)

	assert.Equal(t, got, want)
}

func TestPartTwo(t *testing.T) {
	input := file.MustReadToString("test.txt")
	want := 25272
	solver := Solver{}
	solver.Init(true)

	got := solver.Part2(input)

	assert.Equal(t, got, want)
}

func TestParseLine(t *testing.T) {
	input := "162,817,812"
	id := 123
	want := junctionBox{
		id:        id,
		x:         162,
		y:         817,
		z:         812,
		circuitID: id,
	}

	got := parseLine(input, id)

	assert.DeepEqual(t, got, want)
}

func TestNewPlayground(t *testing.T) {
	input := file.MustReadToString("test.txt")

	playground := newPlayground(input, logger.New(true))

	assert.Equal(t, len(playground.junctionBoxes), 20)
	assert.Equal(t, len(playground.circuits), 20)
}

func TestJoinCircuits(t *testing.T) {
	playground := playground{
		circuits: map[int][]int{
			0: {0, 1},
			1: {2, 3, 4},
		},
	}

	want := map[int][]int{
		0: {0, 1, 2, 3, 4},
	}

	playground.joinCircuits(0, 1)

	assert.DeepEqual(t, playground.circuits, want)
}

func TestConnect(t *testing.T) {
	t.Run("Expect same circuit ID to return false and no change to state", func(t *testing.T) {
		initialPlayground := playground{
			junctionBoxes: []junctionBox{
				{circuitID: 0},
				{circuitID: 0},
				{circuitID: 1},
				{circuitID: 1},
				{circuitID: 1},
			},
			circuits: map[int][]int{
				0: {0, 1},
				1: {2, 3, 4},
			},
		}

		initialPlayground.connect(0, 1)

		wantPlayground := playground{
			junctionBoxes: []junctionBox{
				{circuitID: 0},
				{circuitID: 0},
				{circuitID: 1},
				{circuitID: 1},
				{circuitID: 1},
			},
			circuits: map[int][]int{
				0: {0, 1},
				1: {2, 3, 4},
			},
		}

		assert.DeepEqual(t, initialPlayground, wantPlayground)
	})

	t.Run("Expect all boxes on joined circuit to use new circuit ID", func(t *testing.T) {
		initialPlayground := playground{
			junctionBoxes: []junctionBox{
				{circuitID: 0},
				{circuitID: 0},
				{circuitID: 1},
				{circuitID: 1},
				{circuitID: 1},
			},
			circuits: map[int][]int{
				0: {0, 1},
				1: {2, 3, 4},
			},
		}

		initialPlayground.connect(1, 2)

		wantPlayground := playground{
			junctionBoxes: []junctionBox{
				{circuitID: 0},
				{circuitID: 0},
				{circuitID: 0},
				{circuitID: 0},
				{circuitID: 0},
			},
			circuits: map[int][]int{
				0: {0, 1, 2, 3, 4},
			},
		}

		assert.DeepEqual(t, initialPlayground, wantPlayground)
	})
}
