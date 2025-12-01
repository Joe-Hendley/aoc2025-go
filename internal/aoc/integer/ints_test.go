package integer

import (
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
