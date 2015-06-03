package xstrings

import (
	"strings"
)

func Start_with(s string, prefix string) bool {

	return strings.HasPrefix(s, prefix)
}
