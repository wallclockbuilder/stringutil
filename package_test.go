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
}
