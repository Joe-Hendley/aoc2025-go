package p7

import (
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	lines := strings.Split(input, "\n")

	row := 0
	beams := make([]bool, len(lines[0]))
	splitterPositions := [][]int{}

	for _, line := range lines {
		splittersOnLine := []int{}
		for charIndex, char := range line {
			switch char {
			case 'S':
				beams[charIndex] = true
			case '^':
				splittersOnLine = append(splittersOnLine, charIndex)
			}
		}
		splitterPositions = append(splitterPositions, splittersOnLine)
	}

	splits := 0

	for row < len(lines) {
		for _, splitterIndex := range splitterPositions[row] {
			if beams[splitterIndex] {
				beams[splitterIndex] = false
				beams[splitterIndex+1] = true
				beams[splitterIndex-1] = true
				splits++
			}
		}
		row++
	}

	return splits
}

type splitter struct {
	position  int
	timelines int
}

func (s *Solver) Part2(input string) int {
	lines := strings.Split(input, "\n")

	row := len(lines) - 1
	splitters := [][]splitter{}

	for _, line := range lines {
		splittersOnLine := []splitter{}
		for charIndex, char := range line {
			if char == '^' {
				splittersOnLine = append(splittersOnLine, splitter{position: charIndex, timelines: 0})
			}
		}
		splitters = append(splitters, splittersOnLine)
	}

	for row >= 0 {
		for splitterIndex, splitter := range splitters[row] {
			lookback := func(col int) {
				lookbackRow := row
				// ugly but it's advent of code
				for {
					if lookbackRow >= len(lines) {
						splitters[row][splitterIndex].timelines++
						return
					}

					for _, lookbackSplitter := range splitters[lookbackRow] {
						if lookbackSplitter.position == col {
							splitters[row][splitterIndex].timelines += lookbackSplitter.timelines
							return
						}
					}
					lookbackRow++
				}
			}
			lookback(splitter.position - 1)
			lookback(splitter.position + 1)
		}
		row--
	}

	// return the timelines from the first splitter
	for row := range lines {
		for _, splitter := range splitters[row] {
			return splitter.timelines
		}
	}
	return 0
}
