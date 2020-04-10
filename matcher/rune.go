package matcher

import (
	"github.com/viant/parsly"
	"unicode/utf8"
)

type Rune struct {
	value rune
}

func (c *Rune) Match(cursor *parsly.Cursor) (matched int) {
	runeValue, width := utf8.DecodeRune(cursor.Input[cursor.Pos:])
	if c.value == runeValue {
		return width
	}
	return 0
}

//NewRune creates a rune matcher
func NewRune(value rune) *Rune {
	return &Rune{value: value}
}
