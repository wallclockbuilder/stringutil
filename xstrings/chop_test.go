package xstrings

import (
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func TestChop(t *testing.T) {
	assert.Equal(t, "string", Chop("string\r\n"))
	assert.Equal(t, "string\n", Chop("string\n\r"))
	assert.Equal(t, "string", Chop("string\n"))
	assert.Equal(t, "strin", Chop("string"))
	assert.Equal(t, "", Chop(Chop("x")))
}
