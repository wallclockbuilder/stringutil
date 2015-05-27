package rugo

import (
	"github.com/wallclockbuilder/rugo/xstrings"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestImports(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("ABCDE", xstrings.Capitalize("abcde"))
	assert.Equal("capt. planet", xstrings.Prepend("capt. ", "planet"))
	assert.Equal("strin", xstrings.Chop("string"))
	assert.Equal("hello", xstrings.Chomp("hello", ""))
	assert.Equal("he", xstrings.Chomp("hello", "llo"))
	assert.Equal("a", xstrings.Chr("abcde"))
	assert.Equal("desserts", xstrings.Reverse("stressed"))
}
