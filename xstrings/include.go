package xstrings

import (
	"strings"
)

// Include returns true if str contains the given string
// or character.
func Include(s string, substr string) bool {
	return strings.Contains(s, substr)
}
