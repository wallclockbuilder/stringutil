package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestLength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, Length("hello"))
}
