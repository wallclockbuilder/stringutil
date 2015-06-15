package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleCapitalize(){
	fmt.Println(Capitalize("bare necessities"))
	// Output: BARE NECESSITIES
}

func TestCapitalize(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("ABCDE", Capitalize("abcde"))
}
