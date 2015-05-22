package xstrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepend(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("capt. planet", Prepend("capt. ", "planet"))
}
