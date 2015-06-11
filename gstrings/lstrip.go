package gstrings

import (
	"strings"
	"unicode"
)

// Lstrip returns a copy of s
// with leading whitespace removed.
// See also gstrings.Rstrip and gstrings.Strip
func Lstrip(s string) string {
	return strings.TrimLeftFunc(s, f)
}

func f(r rune) bool {
	return unicode.IsSpace(r)
}
