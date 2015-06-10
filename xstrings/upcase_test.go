package xstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleUpcase() {
	fmt.Println(Upcase("hEllO"))
	// Output: HELLO
}

func TestUpcase(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("HELLO", Upcase("hEllO"))
}
