package p9

import (
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/must"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/numeric"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	redTiles := []pair{}
	for _, line := range strings.Split(input, "\n") {
		ints := must.StringSplitToInts(line, ",")
		redTiles = append(redTiles, pair{
			x: ints[0],
			y: ints[1],
		})
	}

	largestArea := 0

	for tile1Index := range redTiles {
		for tile2Index := tile1Index + 1; tile2Index < len(redTiles); tile2Index++ {
			area := numeric.IntArea(
				redTiles[tile1Index].x, redTiles[tile1Index].y,
				redTiles[tile2Index].x, redTiles[tile2Index].y,
			)

			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}

type pair struct {
	x, y int
}

func (s *Solver) Part2(input string) int {
	return 0
}
