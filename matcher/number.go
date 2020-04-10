package matcher

import (
	"github.com/viant/parsly"
)

type number struct{}

//TokenMatch matches a number
func (n *number) Match(cursor *parsly.Cursor) (matched int) {
	input := cursor.Input
	pos := cursor.Pos
	if isSing := input[pos] == '-'; isSing {
		pos++
	}
	size := len(input)
	hasDecPoint := false
	hasExponent := false
	valid := false
	var i int
outer:
	for i = pos; i < size; i++ {
		switch input[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			valid = true
		case 'e', 'E':
			if !valid || hasExponent {
				return 0
			}
			hasExponent = true
			if i+1 < size {
				switch input[i+1] {
				case '+', '-':
					i++
				}
			}
			valid = false
		case '.':
			if !valid || hasDecPoint {
				return 0
			}
			valid = false
			hasDecPoint = true
		default:
			break outer
		}
	}

	if !valid {
		return 0
	}
	return i - cursor.Pos
}

//NewNumber creates a number matcher
func NewNumber() *number {
	return &number{}
}
