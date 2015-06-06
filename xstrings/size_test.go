package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestSize(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5, Size("hello"))
}
