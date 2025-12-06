package p6

import (
	"fmt"
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/fun"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/must"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func solveColumnPartOne(col []string) int {
	ints := fun.Map(col[:len(col)-1], func(s string) int {
		return must.Atoi(s)
	})

	operator := col[len(col)-1]

	result := 0
	switch operator {
	case "+":
		fun.Apply(ints, func(i int) {
			result += i
		})
	case "*":
		result++
		fun.Apply(ints, func(i int) {
			result *= i
		})
	default:
		panic(fmt.Sprintf("invalid operation: %s", operator))
	}

	return result
}

func (s *Solver) Part1(input string) int {
	lines := strings.Split(input, "\n")
	fields := fun.Map(lines, func(s string) []string {
		return strings.Fields(s)
	})

	result := 0

	for colIndex := range len(fields[0]) {
		col := []string{}

		for lineIndex := range len(fields) {
			col = append(col, fields[lineIndex][colIndex])
		}

		result += solveColumnPartOne(col)
	}

	return result
}

func (s *Solver) Part2(input string) int {
	parsedInput := parseInputPreservingSpaces(input)

	result := 0

	for colIndex := range len(parsedInput[0]) {
		col := []string{}

		for lineIndex := range len(parsedInput) {
			col = append(col, parsedInput[lineIndex][colIndex])
		}

		result += solveColumnPartTwo(col)
	}

	return result
}

func parseInputPreservingSpaces(input string) [][]string {
	lines := strings.Split(input, "\n")
	operatorLine := strings.Fields(lines[len(lines)-1])
	intLines := lines[:len(lines)-1]

	// work out how long each field is before we try to split
	fields := fun.Map(intLines, func(s string) []string {
		return strings.Fields(s)
	})

	split := make([][]string, len(intLines))
	for i := range split {
		split[i] = []string{}
	}

	colPointer := 0
	for colIndex := range len(fields[0]) {
		colSize := 0
		for lineIndex := range len(fields) {
			fieldSize := len(fields[lineIndex][colIndex])
			if fieldSize > colSize {
				colSize = fieldSize
			}
		}

		for lineIndex := range len(fields) {
			split[lineIndex] = append(split[lineIndex], intLines[lineIndex][colPointer:colPointer+colSize])
		}

		colPointer += colSize + 1
	}

	split = append(split, operatorLine)

	return split
}

func solveColumnPartTwo(col []string) int {
	ints := transformColToWeirdOctopusNumbers(col[:len(col)-1])

	operator := col[len(col)-1]

	result := 0
	switch operator {
	case "+":
		fun.Apply(ints, func(i int) {
			result += i
		})
	case "*":
		result++
		fun.Apply(ints, func(i int) {
			result *= i
		})
	default:
		panic(fmt.Sprintf("invalid operation: %s", operator))
	}

	return result
}

func transformColToWeirdOctopusNumbers(col []string) []int {
	i := 0
	ints := []int{}

	for {
		foundDigit := false
		digits := []byte{}
		for _, line := range col {
			if i < len(line) && line[i] != ' ' {
				foundDigit = true
				digits = append(digits, line[i])
			}
		}

		if !foundDigit {
			break
		}

		ints = append(ints, must.Atoi(string(digits)))
		i++
	}

	return ints
}
