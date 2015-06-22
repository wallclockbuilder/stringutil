package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleInclude(){
	fmt.Println(Include("hello", "lo"))
	fmt.Println(Include("hello", "ol"))
	fmt.Println(Include("hello", "?h"))
	// Output: true
	// false
	// true
}

func TestInclude(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, Include("hello", "lo"))
	assert.Equal(false, Include("hello", "ol"))
	assert.Equal(true, Include("hello", "?h"))
}
