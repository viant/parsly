package matcher

import (
	"github.com/viant/parsly"
	"unicode/utf8"
)

type Runes struct {
	values []rune
	size   int
}

func (c *Runes) Match(cursor *parsly.Cursor) (matched int) {
	runeValue, width := utf8.DecodeRune(cursor.Input[cursor.Pos:])
	i := 0
loop:
	if runeValue == c.values[i] {
		return width
	}
	i++
	if i < c.size {
		goto loop
	}
	return 0
}

//NewRunes creates a runes matcher
func NewRunes(values []rune) *Runes {
	return &Runes{values: values, size: len(values)}
}
