package grid

import (
	"iter"
	"strings"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/grid/direction"
)

type Grid[T comparable] struct {
	inner  []T
	width  int
	height int
}

func New[T comparable](width, height int, defaultValue T) Grid[T] {
	inner := make([]T, width*height)
	for i := range inner {
		inner[i] = defaultValue
	}

	return Grid[T]{
		inner:  inner,
		width:  width,
		height: height,
	}
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

type GridItem[T comparable] struct {
	x, y  int
	value T
}

func (gi GridItem[T]) X() int {
	return gi.x
}

func (gi GridItem[T]) Y() int {
	return gi.y
}

func (gi GridItem[T]) Item() T {
	return gi.value
}

func (g Grid[T]) all() []GridItem[T] {
	x := 0
	y := 0

	items := make([]GridItem[T], 0, len(g.inner))

	for _, v := range g.inner {
		if x >= g.Width() {
			x = 0
			y++
		}

		items = append(items, GridItem[T]{x: x, y: y, value: v})

		x++
	}

	return items
}

func (g Grid[T]) All() iter.Seq[GridItem[T]] {
	return func(yield func(GridItem[T]) bool) {
		for _, v := range g.all() {
			if !yield(v) {
				return
			}
		}
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

func (g Grid[T]) Replace(x, y int, value T) {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		return
	}

	g.inner[g.idx(x, y)] = value
}

func (g Grid[T]) MapToNeighbours(x, y int, f func(x, y int, value T)) {
	for _, d := range direction.All() {
		dx := x + d.X()
		dy := y + d.Y()
		if g.InBounds(dx, dy) {
			f(dx, dy, g.At(dx, dy))
		}
	}
}
