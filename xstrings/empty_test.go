package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestEmpty(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(false, Empty("hello"))
	assert.Equal(false, Empty(" "))
	assert.Equal(true, Empty(""))
}
