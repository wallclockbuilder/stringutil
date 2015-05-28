package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestStrip(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", Strip("    hello     "))
	assert.Equal("goodbye", Strip("\tgoodbye\r\n"))
}
