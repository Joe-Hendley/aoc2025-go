package numeric

import "math"

// TODO - number type

func SumInts(ints ...int) int {
	s := 0

	for i := range ints {
		s += ints[i]
	}

	return s
}

func AreIntsEqual(target []int) bool {
	if len(target) > 2 {
		for i := range target {
			if target[i] != target[0] {
				return false
			}
		}
	}

	return true
}

func AreIntsEqualTo(target []int, n int) bool {
	if len(target) > 2 {
		for i := range target {
			if target[i] != n {
				return false
			}
		}
	}

	return true
}

func LCM(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	gcd := euclid(nums[0], nums[1])

	if nums[0] < 0 {
		nums[0] *= -1
	}

	if nums[1] < 0 {
		nums[1] *= -1
	}

	lcm := (nums[0] * nums[1]) / gcd

	return LCM(append(nums[2:], lcm))
}

func euclid(a, b int) int {
	c := 0
	for a != b {
		if a > b {
			c = a - b
			a = c
		} else {
			c = b - a
			b = c
		}
	}

	return a
}

func IntDistance1D(a, b int) int {
	x := a - b
	if x > 0 {
		return x
	}
	return -x
}

func IntArea(x1, y1, x2, y2 int) int {
	dX := IntDistance1D(x1, x2) + 1
	dY := IntDistance1D(y1, y2) + 1
	return dX * dY
}

func IntDistance2D(x1, x2, y1, y2 int) float64 {
	return Distance2D(
		float64(x1), float64(x1),
		float64(y1), float64(y1),
	)

}

func Distance2D(x1, x2, y1, y2 float64) float64 {
	xComponent := math.Pow(x1-x2, 2)
	yComponent := math.Pow(y1-y2, 2)
	return math.Sqrt(xComponent + yComponent)
}

func IntDistance3D(x1, x2, y1, y2, z1, z2 int) float64 {
	return Distance3D(
		float64(x1), float64(x2),
		float64(y1), float64(y2),
		float64(z1), float64(z2),
	)
}

func Distance3D(x1, x2, y1, y2, z1, z2 float64) float64 {
	xComponent := math.Pow(x1-x2, 2)
	yComponent := math.Pow(y1-y2, 2)
	zComponent := math.Pow(z1-z2, 2)
	return math.Sqrt(xComponent + yComponent + zComponent)
}
