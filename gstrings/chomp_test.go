package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleChomp() {
	fmt.Println(Chomp("hello", "llo"))
	fmt.Println(Chomp("hello\n", ""))
	fmt.Println(Chomp("hello\r\n", ""))
	// Output: he
	// hello
	// hello
}

func TestChomp(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("hello", Chomp("hello", ""))
	assert.Equal("hello", Chomp("hello\n", ""))
	assert.Equal("hello", Chomp("hello\r\n", ""))
	assert.Equal("hello\n", Chomp("hello\n\r", ""))
	assert.Equal("hello", Chomp("hello\r", ""))
	assert.Equal("hello \n there", Chomp("hello \n there", ""))
	assert.Equal("he", Chomp("hello", "llo"))
	assert.Equal("hello", Chomp("hello\r\n\r\n", ""))
	assert.Equal("hello\r\n\r", Chomp("hello\r\n\r\r\n", ""))
}
