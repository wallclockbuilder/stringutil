package rugo

import (
	"github.com/wallclockbuilder/rugo/gstrings"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestImports(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("ABCDE", gstrings.Capitalize("abcde"))
	assert.Equal("capt. planet", gstrings.Prepend("planet", "capt. "))
	assert.Equal("strin", gstrings.Chop("string"))

	//Chomp()
	assert.Equal("hello", gstrings.Chomp("hello", ""))
	assert.Equal("he", gstrings.Chomp("hello", "llo"))
	assert.Equal("a", gstrings.Chr("abcde"))
	assert.Equal("desserts", gstrings.Reverse("stressed"))

	assert.Equal("hello", gstrings.Strip("    hello     "))

	assert.Equal("hello", gstrings.Downcase("hEllO"))

	assert.Equal("HELLO", gstrings.Upcase("hEllO"))

	assert.Equal(5, gstrings.Length("hello"))

	assert.Equal(true, gstrings.Include("hello", "lo"))

	assert.Equal(false, gstrings.Empty("hello"))

	assert.Equal(true, gstrings.Start_with("hello", "hell"))

	assert.Equal("hello   ", gstrings.Lstrip("   hello   "))

	assert.Equal("   hello", gstrings.Rstrip("   hello   "))

	assert.Equal(5, gstrings.Size("hello"))
}
