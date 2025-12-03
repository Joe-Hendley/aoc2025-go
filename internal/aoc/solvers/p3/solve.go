package p3

import (
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/fun"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/integer"
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

	return integer.Sum(fun.Map(lines, solveLinePartOne)...)
}

func solveLinePartOne(input string) int {
	return solveLine(input, 2)
}

func (s *Solver) Part2(input string) int {
	lines := strings.Split(input, "\n")

	return integer.Sum(fun.Map(lines, solveLinePartTwo)...)
}

func solveLinePartTwo(line string) int {
	return solveLine(line, 12)
}

func solveLine(line string, numDigits int) int {
	digits := make([]byte, numDigits)
	leftBound := 0
	rightBound := len(line) - numDigits + 1

	for digitIndex := range digits {
		maxDigitIndex := leftBound
		for pointerIndex := leftBound; pointerIndex < rightBound; pointerIndex++ {
			if line[pointerIndex] > digits[digitIndex] {
				maxDigitIndex = pointerIndex
				digits[digitIndex] = line[pointerIndex]
			}
		}

		leftBound = maxDigitIndex + 1
		rightBound = rightBound + 1
	}

	return must.Atoi(string(digits))
}
