package direction

import "slices"

type Direction struct {
	x, y int
}

func (d Direction) X() int {
	return d.x
}

func (d Direction) Y() int {
	return d.y
}

var (
	UP        = Direction{x: 0, y: -1}
	DOWN      = Direction{x: 0, y: 1}
	LEFT      = Direction{x: -1, y: 0}
	RIGHT     = Direction{x: 1, y: 0}
	UPLEFT    = Direction{x: -1, y: -1}
	UPRIGHT   = Direction{x: 1, y: -1}
	DOWNLEFT  = Direction{x: -1, y: 1}
	DOWNRIGHT = Direction{x: 1, y: 1}
)

func Cardinal() []Direction {
	return []Direction{
		UP, DOWN, LEFT, RIGHT,
	}
}

func Diagonal() []Direction {
	return []Direction{
		UPLEFT, UPRIGHT, DOWNLEFT, DOWNRIGHT,
	}
}

func All() []Direction {
	return slices.Concat(Cardinal(), Diagonal())
}

func (d Direction) Opposite() Direction {
	return Direction{
		x: d.x * -1,
		y: d.y * -1,
	}
}
