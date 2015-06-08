package xstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleDowncase() {
	fmt.Println(Downcase("hEllO"))
	// Output: hello
}

func TestDowncase(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", Downcase("hEllO"))
}
