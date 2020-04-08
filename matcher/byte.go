package matcher

import (
	"github.com/viant/parsly"
)

type Byte struct {
	value byte
}

func (c *Byte) Match(cursor *parsly.Cursor) (matched int) {
	if  cursor.Input[cursor.Pos] == c.value {
		return 1
	}
	return 0
}

//NewByte creates a rune matcher
func NewByte(value byte) *Byte {
	return &Byte{value: value}
}


