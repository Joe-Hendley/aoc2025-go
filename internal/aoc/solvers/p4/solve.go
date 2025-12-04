package p4

import (
	"github.com/Joe-Hendley/aoc2025/internal/aoc/grid"
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
	paperStore := grid.FromString(input)

	accessibleRolls := 0

	for gridItem := range paperStore.All() {
		if gridItem.Item() == paperRoll {
			neighbours := 0
			paperStore.MapToNeighbours(gridItem.X(), gridItem.Y(), func(_, _ int, value rune) {
				if value == paperRoll {
					neighbours++
				}
			})

			if neighbours < 4 {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func (s *Solver) Part2(input string) int {
	paperStore := grid.FromString(input)
	neighboursGrid := grid.New(paperStore.Width(), paperStore.Height(), 0)

	accessibleRolls := []struct{ x, y int }{}

	for gridItem := range paperStore.All() {
		if gridItem.Item() == paperRoll {
			neighbours := 0
			paperStore.MapToNeighbours(gridItem.X(), gridItem.Y(), func(_, _ int, value rune) {
				if value == paperRoll {
					neighbours++
				}
			})

			neighboursGrid.Replace(gridItem.X(), gridItem.Y(), neighbours)

			if neighbours < 4 {
				accessibleRolls = append(accessibleRolls, struct {
					x int
					y int
				}{gridItem.X(), gridItem.Y()})
			}
		}
	}

	removedRolls := 0

	for len(accessibleRolls) > 0 {
		roll := accessibleRolls[0]

		if paperStore.At(roll.x, roll.y) == paperRoll {
			paperStore.MapToNeighbours(roll.x, roll.y, func(x, y int, value rune) {
				if value != paperRoll {
					return
				}

				neighboursGrid.Replace(x, y, neighboursGrid.At(x, y)-1)
				if neighboursGrid.At(x, y) < 4 {
					accessibleRolls = append(accessibleRolls, struct {
						x int
						y int
					}{x, y})
				}
			})

			paperStore.Replace(roll.x, roll.y, emptySpace)
			removedRolls++
		}

		accessibleRolls = accessibleRolls[1:]
	}

	return removedRolls
}
