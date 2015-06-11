package gstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestInclude(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, Include("hello", "lo"))
	assert.Equal(false, Include("hello", "ol"))
	assert.Equal(true, Include("hello", "?h"))
}
