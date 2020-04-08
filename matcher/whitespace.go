package matcher

import "github.com/viant/parsly"

type Whitespace struct{}

//TokenMatch matches whitespaces
func (w *Whitespace) Match(cursor *parsly.Cursor) (matched int) {
	offset := cursor.Pos
	input := cursor.Input
	size := cursor.InputSize
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

//NewWhiteSpace creates a Whitespace matcher
func NewWhiteSpace() *Whitespace {
	return &Whitespace{}
}
