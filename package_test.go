package rugo

import (
	"github.com/wallclockbuilder/rugo/xstrings"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestImports(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("ABCDE", xstrings.Capitalize("abcde"))
	assert.Equal("capt. planet", xstrings.Prepend("planet", "capt. "))
	assert.Equal("strin", xstrings.Chop("string"))

	//Chomp()
	assert.Equal("hello", xstrings.Chomp("hello", ""))
	assert.Equal("he", xstrings.Chomp("hello", "llo"))
	assert.Equal("a", xstrings.Chr("abcde"))
	assert.Equal("desserts", xstrings.Reverse("stressed"))

	assert.Equal("hello", xstrings.Strip("    hello     "))

	assert.Equal("hello", xstrings.Downcase("hEllO"))

	assert.Equal("HELLO", xstrings.Upcase("hEllO"))

	assert.Equal(5, xstrings.Length("hello"))

	assert.Equal(true, xstrings.Include("hello", "lo"))

	assert.Equal(false, xstrings.Empty("hello"))

	assert.Equal(true, xstrings.Start_with("hello", "hell"))

	assert.Equal("hello   ", xstrings.Lstrip("   hello   "))
}
