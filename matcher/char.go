package matcher

import (
	"github.com/viant/parsly"
)


//NewChar creates a rune or a byte matcher
func NewChar(value rune) parsly.Matcher {
	if isByte(value) {
		return NewByte(byte(value))
	}
	return NewRune(value)
}


