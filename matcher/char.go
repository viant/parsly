package matcher

import (
	"github.com/viant/parsly/lex"
	"unicode/utf8"
)

type utfChar struct {
	value rune
}

func (c *utfChar) Match(input []byte, offset int) (matched int) {
	runeValue, width := utf8.DecodeRune(input[offset:])
	if c.value == runeValue {
		return width
	}
	return 0
}


type char struct {
	value byte
}

func (c *char) Match(input []byte, offset int) (matched int) {
	runeValue := input[offset]
	if c.value == runeValue {
		return 1
	}
	return 0
}

//MatchRune matches rune
func (c *char) MatchRune(runeValue rune) bool {
	return byte(runeValue) == c.value
}

//NewChar creates a rune matcher
func NewChar(value rune) lex.Matcher {
	if isByte(value) {
		return &char{value: byte(value)}
	}
	return &utfChar{value: value}
}
