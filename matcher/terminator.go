package matcher

import (
	"github.com/viant/parsly"
)

type Terminator struct {
	value     byte
	inclusive bool
}

func (t *Terminator) Match(cursor *parsly.Cursor) (matched int) {
	hasMatch := false
	for _, c := range cursor.Input[cursor.Pos:] {
		matched++
		if hasMatch = c == t.value; hasMatch {
			if !t.inclusive {
				matched--
			}
			break
		}
	}
	if !hasMatch {
		return 0
	}
	return matched
}

//Terminator creates a terminator byte matcher
func NewTerminator(value byte, inclusive bool) *Terminator {
	return &Terminator{value: value, inclusive: inclusive}
}
