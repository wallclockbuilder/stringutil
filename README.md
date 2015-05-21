Rugo
 ====
Ruby's string manipulation magic brought to Golang

> **Why?**
>  "What was the biggest surprise you encountered rolling out Go?" I knew the answer instantly: Although we expected C++ programmers to see Go as an alternative, instead **most Go programmers come from languages like Python and Ruby**. 
 Robert Pike, [Less is exponentially more](http://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html)


Rugo has made all the string manipulation methods from ruby accessible in go.


Install
---------
<code>
go get https://github.com/wallclockbuilder/rugo


Import
----------
<code>
import "xstrings" "https://github.com/wallclockbuilder/rugo"


Use
-----
<code>
xstrings.Capitalize("abcde") 		  	      #=> 		"ABCDE"

xstrings.Prepend("captain " , "planet") 	#=> 		"captain planet"

xstrings.Swapcase("Hello") 		          	#=> 		"hELLO"
</code>
