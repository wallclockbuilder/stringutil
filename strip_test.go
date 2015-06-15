package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleStrip(){
	fmt.Println(Strip("   hello   "))
	fmt.Println(Strip("\tgoodbye\r\n"))
	// Output: hello
	// goodbye
}

func TestStrip(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", Strip("    hello     "))
	assert.Equal("goodbye", Strip("\tgoodbye\r\n"))
}
