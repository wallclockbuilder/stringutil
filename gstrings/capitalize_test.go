package gstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestCapitalize(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("ABCDE", Capitalize("abcde"))
}
