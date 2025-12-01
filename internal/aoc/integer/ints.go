package integer

func Sum(ints ...int) int {
	s := 0

	for i := range ints {
		s += ints[i]
	}

	return s
}

func IsEqual(target []int) bool {
	if len(target) > 2 {
		for i := range target {
			if target[i] != target[0] {
				return false
			}
		}
	}

	return true
}

func IsEqualTo(target []int, n int) bool {
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

func Distance(a, b int) int {
	x := a - b
	if x > 0 {
		return x
	}
	return -x
}
