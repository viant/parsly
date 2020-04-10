package example

import (
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher"
)

var Whitespace = parsly.NewToken(0, "Whitespace", matcher.NewWhiteSpace())
var Number = parsly.NewToken(1, "Number", matcher.NewNumber())
var Term = parsly.NewToken(2, "Term", matcher.NewCharset("+-"))
var Factor = parsly.NewToken(3, "Factor", matcher.NewCharset("*/"))
