package matcher

import (
	"github.com/viant/parsly"
)

type Terminator struct {
	value byte
}


func (t *Terminator) Match(cursor *parsly.Cursor) (matched int) {
	hasMatch := false
	for _, c := range cursor.Input[cursor.Pos:] {
		matched++
		if hasMatch = c == t.value; hasMatch {
			break
		}
	}
	if ! hasMatch {
		return 0
	}
	return matched
}


//Terminator creates a terminator byte matcher
func NewTerminator(value byte) *Terminator {
	return &Terminator{value: value}
}
