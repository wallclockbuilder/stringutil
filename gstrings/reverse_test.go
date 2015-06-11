package gstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", Reverse(""))
	assert.Equal("a", Reverse("a"))
	assert.Equal("oy", Reverse("yo"))
	assert.Equal("wow", Reverse("wow"))
	assert.Equal("dog", Reverse("god"))
	assert.Equal("rats", Reverse("star"))
	assert.Equal("desserts", Reverse("stressed"))
}
