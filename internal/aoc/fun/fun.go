package fun

func Map[S ~[]T, T, E any](slice S, function func(T) E) []E {
	mapped := make([]E, len(slice))
	for i := range slice {
		mapped[i] = function(slice[i])
	}

	return mapped
}

func MapInPlace[S ~[]T, T any](slice S, function func(T) T) {
	for i := range slice {
		slice[i] = function(slice[i])
	}
}

func Apply[S ~[]T, T any](slice S, function func(T)) {
	for i := range slice {
		function(slice[i])
	}
}

func Filter[S ~[]T, T any](slice S, predicate func(T) bool) S {
	filtered := make([]T, 0, len(slice))
	for i := range slice {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}

	return filtered
}

func FilterMap[K comparable, T any](m map[K]T, predicate func(k K, v T) bool) map[K]T {
	filtered := make(map[K]T, len(m))
	for k, v := range m {
		if predicate(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

func Fold[T any](slice []T, pairFunction func(a, b T) T, zero T) T {
	if len(slice) == 0 {
		return zero
	}

	return pairFunction(slice[0], Fold(slice[1:], pairFunction, zero))
}

type Enumurated[T any] struct {
	Index int
	Value T
}

func Enumerate[S ~[]T, T any](slice S) []Enumurated[T] {
	enumeratedSlice := make([]Enumurated[T], len(slice))
	for i, v := range slice {
		enumeratedSlice[i] = Enumurated[T]{Index: i, Value: v}
	}

	return enumeratedSlice
}

func FlatMap[S ~[]T, T, E any](slice S, function func(T) []E) []E {
	mapped := Map(slice, function)

	flattened := []E{}
	Apply(mapped, func(slice []E) {
		Apply(slice, func(element E) {
			flattened = append(flattened, element)
		})
	})

	return flattened
}
