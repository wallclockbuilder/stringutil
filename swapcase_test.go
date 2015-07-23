package stringutil

import (
  "fmt"
  "testing"
  "github.com/wallclockbuilder/testify/assert"
)

func ExampleSwapcase(){
  fmt.Println(Swapcase("Hello"))
  fmt.Println(Swapcase("cYbEr_PuNk11"))
  fmt.Println(Swapcase("abcdefghijklmnopqrstuvwxyz"))
  fmt.Println(Swapcase("developer HAPPINESS"))
  // Output: hELLO
  // CyBeR_pUnK11
  // ABCDEFGHIJKLMNOPQRSTUVWXYZ
  // DEVELOPER happiness
}

func TestSwapcase(t *testing.T){
  assert.Equal(t, "hELLO", Swapcase("Hello"))
  assert.Equal(t, "CyBeR_pUnK11", Swapcase("cYbEr_PuNk11"))
  assert.Equal(t, "DEVELOPER happiness", Swapcase("developer HAPPINESS"))
  assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", Swapcase("abcdefghijklmnopqrstuvwxyz"))
  assert.Equal(t, "1234567890", Swapcase("1234567890"))
}
