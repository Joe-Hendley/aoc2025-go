package p1

import (
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/must"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	lines := strings.Split(input, "\n")

	position := 50
	result := 0

	for _, line := range lines {
		diff := extractCommand(line)
		position += diff

		if position%100 == 0 {
			position = 0
			result++
		}
	}

	return result
}

func (s *Solver) Part2(input string) int {
	lines := strings.Split(input, "\n")

	position := 50
	result := 0

	for _, line := range lines {
		diff := extractCommand(line)

		result += countClicks(position, diff)

		position += diff

		position = getBoundedPosition(position)
	}

	return result
}

func extractCommand(s string) int {
	if len(s) < 2 {
		return 0
	}

	magnitude := must.Atoi(s[1:])

	switch s[0] {
	case 'R':
		return magnitude
	case 'L':
		return -magnitude
	default:
		return 0
	}
}

func getBoundedPosition(position int) int {
	for position < 0 {
		position += 100
	}

	return position % 100
}

func countClicks(position, diff int) int {
	switch {
	case diff > 0:
		return countClicksPositiveDiff(position, diff)
	case diff < 0:
		return countClicksNegativeDiff(position, diff)
	default:
		return 0
	}
}

func countClicksPositiveDiff(position, diff int) int {
	result := 0
	for diff > 99 {
		result++
		diff -= 100
	}

	if (position + diff) > 99 {
		result++
	}

	return result
}

func countClicksNegativeDiff(position, diff int) int {
	result := 0
	for diff < -99 {
		result++
		diff += 100
	}

	if position != 0 && (position+diff) < 1 {
		result++
	}

	return result
}
