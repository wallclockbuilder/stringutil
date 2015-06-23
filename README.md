#gstrings [![Build Status](https://travis-ci.org/wallclockbuilder/gstrings.svg)](https://travis-ci.org/wallclockbuilder/gstrings) [![GoDoc](https://godoc.org/github.com/wallclockbuilder/gstrings?status.svg)](http://godoc.org/github.com/wallclockbuilder/gstrings)

Ruby's string manipulation magic brought to Golang

Why?
>  "What was the biggest surprise you encountered rolling out Go?" I knew the answer instantly: Although we expected C++ programmers to see Go as an alternative, instead **most Go programmers come from languages like Python and Ruby**.
 Robert Pike, [Less is exponentially more](http://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html)

gstrings makes all the string manipulation methods from ruby accessible in go.


###Install
```bash
 go get https://github.com/wallclockbuilder/gstrings
```

###Import
```golang
 import "gstrings" "https://github.com/wallclockbuilder/gstrings"
```
### Example
I want to remove the trailing space
 from the end of a string.

This is what Go forces me to do:
```golang
strings.TrimRightFunc(s, unicode.IsSpace)
```
This is the simple version from Ruby:
```golang
gstrings.Rstrip(s)
```
I prefer the Ruby version.

###Use
```golang
package main
import "github.com/wallclockbuilder/gstrings"

func main() {
  gstrings.Capitalize("abcde")          #=> "ABCDE"
  gstrings.Reverse("stressed")          #=> "desserts"
  gstrings.Swapcase("Hello")            #=> "hELLO"
}
```

###Documentation
`go doc` style documentation for this package is available online at http://godoc.org/github.com/wallclockbuilder/gstrings
###Method list
- [ ] ascii_only?
- [ ] b
- [ ] bytes
- [ ] bytesize
- [ ] byteslice
- [x] capitalize
- [ ] capitalize!
- [ ] casecmp
- [ ] center
- [ ] chars
- [x] chomp
- [ ] chomp!
- [x] chop
- [ ] chop!
- [x] chr
- [ ] clear
- [ ] codepoints
- [ ] concat
- [ ] count
- [ ] crypt
- [ ] delete
- [ ] delete!
- [x] downcase
- [ ] downcase!
- [ ] dump
- [ ] each_byte
- [ ] each_char
- [ ] each_codepoint
- [ ] each_line
- [x] empty?
- [ ] encode
- [ ] encode!
- [ ] encoding
- [ ] end_with?
- [ ] eql?
- [ ] force_encoding
- [ ] getbyte
- [ ] gsub
- [ ] gsub!
- [ ] hash
- [ ] hex
- [x] include?
- [ ] index
- [ ] initialize_copy
- [ ] insert
- [ ] inspect
- [ ] intern
- [x] length
- [ ] lines
- [ ] ljust
- [x] lstrip
- [ ] lstrip!
- [ ] match
- [ ] next
- [ ] next!
- [ ] oct
- [ ] ord
- [ ] partition
- [x] prepend
- [ ] replace
- [x] reverse
- [ ] reverse!
- [ ] rindex
- [ ] rjust
- [ ] rpartition
- [x] rstrip
- [ ] rstrip!
- [ ] scan
- [ ] scrub
- [ ] scrub!
- [ ] setbyte
- [x] size
- [ ] slice
- [ ] slice!
- [ ] split
- [ ] squeeze
- [ ] squeeze!
- [x] start_with?
- [x] strip
- [ ] strip!
- [ ] sub
- [ ] sub!
- [ ] succ
- [ ] succ!
- [ ] sum
- [ ] swapcase
- [ ] swapcase!
- [ ] to_c
- [ ] to_f
- [ ] to_i
- [ ] to_r
- [ ] to_s
- [ ] to_str
- [ ] to_sym
- [ ] tr
- [ ] tr!
- [ ] tr_s
- [ ] tr_s!
- [ ] unpack
- [x] upcase
- [ ] upcase!
- [ ] upto
- [ ] valid_encoding?

## Contributing

1. Fork it ( https://github.com/wallclockbuilder/gstrings )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
