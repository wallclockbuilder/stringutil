package rugo

import (
	"github.com/stretchr/testify/assert"
	"github.com/wallclockbuilder/rugo/xstrings"
	"testing"
)

func TestImports(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("ABCDE", xstrings.Capitalize("abcde"))
}
