package p2

import (
	"strconv"
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

func (s *Solver) Part1(input string) int {
	idRanges := parseInput(input)
	sumInvalidIDs := fun.Fold(fun.Map(idRanges, func(ir idRange) int {
		return fun.Fold(ir.findInvalidIDs(), func(a, b int) int {
			return a + b
		}, 0)
	}), func(a, b int) int {
		return a + b
	}, 0)
	return sumInvalidIDs
}

func (s *Solver) Part2(input string) int {
	return 0
}

type idRange struct {
	start int
	end   int
}

func (ir idRange) findInvalidIDs() []int {
	invalidIDs := []int{}
	for id := ir.start; id <= ir.end; id++ {
		if !isValidID(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}
	return invalidIDs
}

func parseInput(input string) []idRange {
	idRanges := []idRange{}
	for _, part := range strings.Split(input, ",") {
		idRanges = append(idRanges, parseRange(part))
	}
	return idRanges
}

func parseRange(input string) idRange {
	parts := strings.Split(input, "-")
	return idRange{
		start: must.Atoi(parts[0]),
		end:   must.Atoi(parts[1]),
	}
}

func isValidID(id int) bool {
	stringed := strconv.Itoa(id)
	if len(stringed)%2 != 0 {
		return true
	}

	halfway := len(stringed) / 2
	for index := 0; index < halfway; index++ {
		if stringed[index] != stringed[index+halfway] {
			return true
		}
	}

	return false
}
