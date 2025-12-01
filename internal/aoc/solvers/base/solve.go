package base

import (
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

func (s *Solver) Part1(input string) int {
	return 0
}

func (s *Solver) Part2(input string) int {
	return 0
}
