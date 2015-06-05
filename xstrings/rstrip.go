package xstrings

import (
	"strings"
	"unicode"
)

// Rstrip returns a copy of s
// with trailing whitespace removed.
// See also xstrings.lstrip and xstrings.strip.
func Rstrip(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}
