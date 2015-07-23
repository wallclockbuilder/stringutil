package stringutil

import (
	"fmt"
	"github.com/wallclockbuilder/testify/assert"
	"testing"
)

func ExampleLength(){
	fmt.Println(Length("hello"))
	// Output: 5
}

func TestLength(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(5, Length("hello"))
}
