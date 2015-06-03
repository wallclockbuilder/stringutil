package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestStart_with(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Start_with("hello", "hell"))
}
