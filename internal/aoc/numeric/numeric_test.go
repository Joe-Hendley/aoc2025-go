package numeric

import (
	"fmt"
	"testing"

	"github.com/Joe-Hendley/aoc2025/internal/aoc/assert"
)

func TestEuclid(t *testing.T) {
	got := euclid(252, 105) // wikipedia baaybeee
	assert.Equal(t, got, 21)
}

func TestReduceLCM(t *testing.T) {
	got := LCM([]int{21, 6})
	assert.Equal(t, got, 42)
}

func TestIntDistance1D(t *testing.T) {
	got := IntDistance1D(2, 9)
	assert.Equal(t, got, 7)
}

func TestIntArea(t *testing.T) {
	for _, testCase := range []struct {
		x1, y1, x2, y2, want int
	}{
		{2, 5, 9, 7, 24},
		{7, 1, 11, 7, 35},
		{7, 3, 2, 3, 6},
		{2, 5, 11, 1, 50},
	} {
		t.Run(fmt.Sprintf("%d,%d %d,%d | %d", testCase.x1, testCase.y1, testCase.x2, testCase.y2, testCase.want), func(t *testing.T) {
			assert.Equal(t, IntArea(testCase.x1, testCase.y1, testCase.x2, testCase.y2), testCase.want)
		})
	}
}
