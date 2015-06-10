package xstrings

import (
	"strings"
	"unicode"
)

// Rstrip returns a copy of s
// with trailing whitespace removed.
// See also xstrings.Lstrip and xstrings.Strip
func Rstrip(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}
