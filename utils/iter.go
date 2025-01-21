package utils

import (
	"iter"
	"strings"
)

// UntilNextEmptyLine returns the line number (1 indexed and the segment)
func UntilNextEmptyLine(s string) iter.Seq2[int, []string] {
	return func(yield func(int, []string) bool) {
		lines := strings.Split(s, "\n")
		previous := 0
		for i, line := range lines {
			if len(strings.TrimSpace(line)) == 0 {
				if !yield(previous+1, lines[previous:i]) {
					return
				}
				previous = i
			}
		}
		yield(previous, lines[previous:])
	}
}

func Zip[T, S any](a []T, b []S) iter.Seq2[T, S] {
	return func(yield func(T, S) bool) {
		for i := 0; i < len(a) && i < len(b); i++ {
			if !yield(a[i], b[i]) {
				return
			}
		}
	}
}

func segmentContainsOnlyBlankLines(segment []string) bool {
	for _, seg := range segment {
		if len(strings.TrimSpace(seg)) != 0 {
			return false
		}
	}
	return true
}
