package iteration

import "strings"

var repeatCount = 3

// Repeat takes an original string and a repeat count and returns the repeated string.
func Repeat(original string, repeatCountArg int) string {
	if repeatCountArg != 0 {
		repeatCount = repeatCountArg
	}

	return strings.Repeat(original, repeatCount)
}
