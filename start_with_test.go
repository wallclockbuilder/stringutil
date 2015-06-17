package gstrings

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleStart_with(){
	fmt.Println(Start_with("hello", "hell"))
	// Output: true
}

func TestStart_with(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, Start_with("hello", "hell"))
}
