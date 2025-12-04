package p4

import (
	"github.com/Joe-Hendley/aoc2025/internal/aoc/grid"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/grid/direction"
	"github.com/Joe-Hendley/aoc2025/internal/aoc/logger"
)

type Solver struct {
	logger.Logger
}

func (s *Solver) Init(verbose bool) {
	s.Logger = logger.New(verbose)
}

const paperRoll rune = '@'
const emptySpace rune = '.'

func (s *Solver) Part1(input string) int {
	paperStack := grid.FromString(input)

	accessibleRolls := 0

	for x := range paperStack.Width() {
		for y := range paperStack.Height() {
			if paperStack.At(x, y) == paperRoll {
				neighbours := 0
				for _, d := range direction.All() {
					if paperStack.CheckCellInDirection(paperRoll, d, x, y) {
						neighbours++
					}
				}

				if neighbours < 4 {
					accessibleRolls++
				}
			}
		}
	}

	return accessibleRolls
}

func (s *Solver) Part2(input string) int {
	paperStack := grid.FromString(input)

	totalAccessibleRolls := 0
	for {
		newPaperStack := grid.New(paperStack.Width(), paperStack.Height(), emptySpace)

		accessibleRolls := 0
		for x := range paperStack.Width() {
			for y := range paperStack.Height() {
				if paperStack.At(x, y) == paperRoll {
					neighbours := 0
					for _, d := range direction.All() {
						if paperStack.CheckCellInDirection(paperRoll, d, x, y) {
							neighbours++
						}
					}

					if neighbours < 4 {
						accessibleRolls++
					} else {
						newPaperStack.Replace(x, y, paperRoll)
					}
				}
			}
		}

		if accessibleRolls == 0 {
			return totalAccessibleRolls
		}

		totalAccessibleRolls += accessibleRolls
		paperStack = newPaperStack
	}
}
