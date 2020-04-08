package matcher

import (
	"github.com/viant/parsly"
)

type Bytes struct {
	values []byte
	size   int
}

func (c *Bytes) Match(cursor *parsly.Cursor) (matched int) {
	b := cursor.Input[cursor.Pos]
	i := 0
loop:
	if b == c.values[i] {
		return 1
	}
	i++
	if i < c.size {
		goto loop
	}
	return 0
}

//NewByte creates a rune matcher
func NewBytes(values []byte) *Bytes {
	return &Bytes{values: values, size: len(values)}
}
