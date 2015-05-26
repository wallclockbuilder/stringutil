package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestChr(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("a", Chr("abcde"))
}
