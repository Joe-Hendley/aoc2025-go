package assert

import "testing"

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func NotEqual[T comparable](t *testing.T, got, notWant T) {
	t.Helper()
	if got == notWant {
		t.Errorf("got %v didn't want %v", got, notWant)
	}
}
