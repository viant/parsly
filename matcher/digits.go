package matcher

import (
	"github.com/viant/parsly"
)

type Digits struct{}

func (d *Digits) Match(cursor *parsly.Cursor) (matched int) {
	size := len(cursor.Input)
outer:
	for i := cursor.Pos; i < size; i++ {
		b := cursor.Input[i]
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			matched++
		default:
			break outer
		}

	}
	return matched
}

//NewDigits creates a Digits matcher
func NewDigits() *Digits {
	return &Digits{}
}
