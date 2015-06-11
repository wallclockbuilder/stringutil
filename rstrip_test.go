package gstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestRstrip(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("   hello", Rstrip("   hello   "))
}
