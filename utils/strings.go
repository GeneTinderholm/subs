package utils

import "strings"

// LeftPad is dangerous, it can take down the whole internet
func LeftPad(s string, length int, padding rune) string {
	if len(s) < length {
		s = strings.Repeat(string(padding), length-len(s)) + s
	}
	return s
}
