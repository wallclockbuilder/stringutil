package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExamplePrepend(){
	fmt.Println(Prepend("planet", "capt. "))
	// Output: capt. planet
}

func TestPrepend(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("capt. planet", Prepend("planet", "capt. "))
}
