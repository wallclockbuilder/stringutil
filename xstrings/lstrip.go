package xstrings

import (
	"strings"
	"unicode"
)

// Lstrip returns a copy of s
// with leading whitespace removed.
// See also xstrings.Rstrip and xstrings.Strip
func Lstrip(s string) string {
	return strings.TrimLeftFunc(s, f)
}

func f(r rune) bool {
	return unicode.IsSpace(r)
}
