package stringutil

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleSize(){
	fmt.Println(Size("hello"))
	// Output: 5
}

func TestSize(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5, Size("hello"))
}
