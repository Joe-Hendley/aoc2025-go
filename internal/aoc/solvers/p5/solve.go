package p5

import (
	"slices"
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
	parts := strings.Split(input, "\n\n")
	freshRanges := parseRanges(s.Logger, parts[0])
	ingredientIDs := must.StringSplitToInts(parts[1], "\n")

	freshIngredients := 0

	fun.Apply(ingredientIDs, func(ingredientID int) {
		if freshRanges.isFresh(ingredientID) {
			freshIngredients++
		}
	})

	return freshIngredients
}

type freshRange struct {
	start, end int
}

func parseRange(input string) freshRange {
	parts := strings.Split(input, "-")
	return freshRange{
		start: must.Atoi(parts[0]),
		end:   must.Atoi(parts[1]),
	}
}

type freshRanges []freshRange

func parseRanges(logger logger.Logger, input string) freshRanges {
	var fr freshRanges = []freshRange{}

	for i, line := range strings.Split(input, "\n") {
		logger.Log("inserting range", line)
		parsedRange := parseRange(line)
		fr = fr.insert(logger, parsedRange)
		logger.Log(i+1, len(fr))
	}

	return fr
}

func (fr freshRanges) insert(logger logger.Logger, newRange freshRange) freshRanges {
	for index, existingRange := range fr {
		switch {
		case newRange.start < existingRange.start && newRange.end < existingRange.start-1:
			if index == 0 {
				logger.Log("inserting at start before", existingRange)
			} else {
				logger.Log("greater than", fr[index-1], "inserting before", existingRange)
			}
			return slices.Insert(fr, index, newRange)
		case newRange.start < existingRange.start && newRange.end <= existingRange.end:
			logger.Log("overlaps start, extending range", existingRange)
			fr[index] = freshRange{start: newRange.start, end: existingRange.end}
			return fr
		case newRange.start <= existingRange.end && newRange.end <= existingRange.end:
			logger.Log("no action, exists in range", existingRange)
			return fr
		case newRange.start <= existingRange.end+1 && newRange.end > existingRange.end:
			if index < len(fr)-1 && newRange.end < fr[index+1].start {
				logger.Log("overlaps end, extending range", existingRange)
				fr[index] = freshRange{start: existingRange.start, end: newRange.end}
				return fr
			} else {
				logger.Log("overlaps end & start of next, extending range", existingRange, fr[index+1])
				fr[index] = freshRange{start: existingRange.start, end: fr[index+1].end}
				return slices.Delete(fr, index+1, index+2)
			}
		case newRange.start > existingRange.end && newRange.end > existingRange.end:
			continue
		default:
			panic("you missed something you doofus")
		}
	}

	if len(fr) == 0 {
		logger.Log("no entries, adding to empty slice")
	} else {
		logger.Log("appending to end after", fr[len(fr)-1])
	}
	appended := append(fr, newRange)

	return appended
}

func (fr freshRanges) isFresh(id int) bool {
	for _, freshRange := range fr {
		if id >= freshRange.start && id <= freshRange.end {
			return true
		}
	}
	return false
}

func (s *Solver) Part2(input string) int {
	parts := strings.Split(input, "\n\n")
	freshRanges := parseRanges(s.Logger, parts[0])

	freshIngredientIDs := 0

	s.Logger.Log(1<<63 - 1)

	for _, freshRange := range freshRanges {
		s.Logger.Logf("%15d - %-15d | %16d | %16d", freshRange.start, freshRange.end, (freshRange.end-freshRange.start)+1, freshIngredientIDs+(freshRange.end-freshRange.start)+1)
		freshIngredientIDs += (freshRange.end - freshRange.start) + 1
	}

	return 0 // until I come back to this properly
	//return freshIngredientIDs
}
