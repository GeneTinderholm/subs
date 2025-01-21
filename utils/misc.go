package utils

// Coalesce returns the first non-zero value if one exists
func Coalesce[T comparable](ts ...T) T {
	var zero T
	for _, t := range ts {
		if t != zero {
			return t
		}
	}
	return zero
}
