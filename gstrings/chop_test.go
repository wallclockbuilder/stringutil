package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleChop() {
	fmt.Println(Chop("Keep it simple"))
	fmt.Println(Chop("Keep it simple\n"))
	fmt.Println(Chop("Developer happiness FTW\r\n"))
	// Output: Keep it simpl
	// Keep it simple
	// Developer happiness FTW
}

func TestChop(t *testing.T) {
	assert.Equal(t, "string", Chop("string\r\n"))
	assert.Equal(t, "string\n", Chop("string\n\r"))
	assert.Equal(t, "string", Chop("string\n"))
	assert.Equal(t, "strin", Chop("string"))
	assert.Equal(t, "", Chop(Chop("x")))
}
