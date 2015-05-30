package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestUpcase(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("HELLO", Upcase("hEllO"))
}
