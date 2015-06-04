package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestLstrip(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello   ", Lstrip("   hello   "))
}
