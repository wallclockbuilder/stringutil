package stringutil

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleChr(){
	fmt.Println(Chr("abcde"))
	// Output: a
}

func TestChr(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("a", Chr("abcde"))
}
