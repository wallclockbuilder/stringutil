package stringutil_test

import (
	"github.com/wallclockbuilder/stringutil"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestImports(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("ABCDE", stringutil.Capitalize("abcde"))
	assert.Equal("capt. planet", stringutil.Prepend("planet", "capt. "))
	assert.Equal("strin", stringutil.Chop("string"))

	//Chomp()
	assert.Equal("hello", stringutil.Chomp("hello", ""))
	assert.Equal("he", stringutil.Chomp("hello", "llo"))
	assert.Equal("a", stringutil.Chr("abcde"))
	assert.Equal("desserts", stringutil.Reverse("stressed"))

	assert.Equal("hello", stringutil.Strip("    hello     "))

	assert.Equal("hello", stringutil.Downcase("hEllO"))

	assert.Equal("HELLO", stringutil.Upcase("hEllO"))

	assert.Equal(5, stringutil.Length("hello"))

	assert.Equal(true, stringutil.Include("hello", "lo"))

	assert.Equal(false, stringutil.Empty("hello"))

	assert.Equal(true, stringutil.Start_with("hello", "hell"))

	assert.Equal("hello   ", stringutil.Lstrip("   hello   "))

	assert.Equal("   hello", stringutil.Rstrip("   hello   "))

	assert.Equal(5, stringutil.Size("hello"))
}
