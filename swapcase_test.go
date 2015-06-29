package gstrings

import (
  "testing"
  "github.com/wallclockbuilder/testify/assert"
)
func TestSwapcase(t *testing.T){
  assert.Equal(t, "hELLO",Swapcase("Hello"))
  assert.Equal(t, "CyBeR_pUnK11", Swapcase("cYbEr_PuNk11"))
  assert.Equal(t, "DEVELOPER happiness", Swapcase("developer HAPPINESS"))
  assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", Swapcase("abcdefghijklmnopqrstuvwxyz"))
  assert.Equal(t, "1234567890", Swapcase("1234567890"))
}
