package matcher

import (
	"github.com/viant/parsly/lex"
)

type whitespace struct{}

//Match matches whitespaces
func (w whitespace) Match(input []byte, offset int) (matched int) {
	size := len(input)
loop:
	{
		index := offset + matched
		if index >= size {
			return matched
		}
		b := input[index]
		if b == ' ' || b == '\n' || b == '\t' || b == '\r' || b == '\v' || b == '\f' || b == 0x85 || b == 0xA0 {
			matched++
			goto loop
		}
	}
	return matched
}


//NewWhiteSpace creates a whitespace matcher
func NewWhiteSpace() lex.Matcher {
	return &whitespace{}
}
