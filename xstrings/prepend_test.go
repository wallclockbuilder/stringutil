package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestPrepend(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("capt. planet", Prepend("planet", "capt. "))
}
