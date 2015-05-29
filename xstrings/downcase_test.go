package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestDowncase(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", Downcase("hEllO"))
}
