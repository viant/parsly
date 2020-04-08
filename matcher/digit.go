package matcher

import (
	"github.com/viant/parsly"
)

type Digit struct{}

func (d *Digit) Match(cursor *parsly.Cursor) (matched int) {
	b := cursor.Input[cursor.Pos]
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 1
	}
	return 0
}

//NewDigit creates a Digit matcher
func NewDigit() *Digit {
	return &Digit{}
}
