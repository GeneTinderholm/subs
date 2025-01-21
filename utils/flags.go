package utils

import "strings"

type Flags map[string]string

// ParseFlags parses flags into a map because the standard parsing interferes with negative numbers
// for the duration (do not use with bool params)
func ParseFlags(args []string) Flags {
	result := Flags{}
	for i := 0; i < len(args)-1; i += 2 {
		if strings.HasPrefix(args[i], "-") {
			result[strings.Trim(args[i], "-")] = args[i+1]
		}
	}
	return result
}
