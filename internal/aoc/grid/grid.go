package grid

import (
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/grid/direction"
)

type Grid[T comparable] struct {
	inner  []T
	width  int
	height int
}

func FromString(input string) Grid[rune] {
	trimmed := strings.TrimSpace(input)
	width := strings.Index(trimmed, "\n")
	singleLine := strings.ReplaceAll(trimmed, "\n", "")
	height := len(singleLine) / width

	return Grid[rune]{
		inner:  []rune(singleLine),
		width:  width,
		height: height,
	}
}

func (g Grid[T]) At(x, y int) T {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		return *new(T)
	}

	return g.inner[g.idx(x, y)]
}

func (g Grid[T]) idx(x, y int) int {
	return (y * g.width) + x
}

func (g Grid[T]) Height() int {
	return g.height
}

func (g Grid[T]) Width() int {
	return g.width
}

func (g Grid[T]) InBounds(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

func (g Grid[T]) CheckCellsInDirection(target []T, d direction.Direction, x, y int) bool {
	targetLength := len(target) - 1 // includes the cell we're on

	if !g.InBounds(
		x+(targetLength*d.X()),
		y+(targetLength*d.Y())) {
		return false
	}

	for i := range target {
		if g.At(x+(d.X()*i), y+(d.Y()*i)) != target[i] {
			return false
		}
	}

	return true
}

func (g Grid[T]) CheckCellInDirection(target T, d direction.Direction, x, y int) bool {
	if !g.InBounds(x+d.X(), y+d.Y()) {
		return false
	}

	return g.At(x+d.X(), y+d.Y()) == target
}
