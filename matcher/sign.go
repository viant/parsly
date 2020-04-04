package matcher

import (
	"github.com/viant/parsly/lex"
)



type sign struct {}

func (c *sign) Match(input []byte, offset int) (matched int) {
	value := input[offset]
	if value == '+' || value == '-' {
		return 1
	}
	return 0
}

//MatchRune matches rune
func (c *sign) MatchRune(value rune) bool {
	if value == '+' || value == '-' {
		return true
	}
	return false
}

//NewSign creates a sign (-/+)  matcher
func NewSign() lex.Matcher {
	return &sign{}
}
