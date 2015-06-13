package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleRstrip(){
	fmt.Println(Rstrip("   hello   "))
	// Output:   hello
}

func TestRstrip(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("   hello", Rstrip("   hello   "))
}
