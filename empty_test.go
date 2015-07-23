package stringutil

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleEmpty(){
	fmt.Println(Empty("hello"))
	fmt.Println(Empty(" "))
	fmt.Println(Empty(""))
	// Output: false
	// false
	// true
}

func TestEmpty(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(false, Empty("hello"))
	assert.Equal(false, Empty(" "))
	assert.Equal(true, Empty(""))
}
