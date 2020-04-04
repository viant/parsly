package matcher

import (
	"github.com/viant/parsly/lex"
	"unicode"
)

type letter struct{}

func (c letter) Match(input []byte, offset int) (matched int) {
	runeValue := rune(input[offset])
	if unicode.IsLetter(runeValue) {
		return 1
	}
	return 0
}

//NewLetter creates a letter matcher
func NewLetter() lex.Matcher {
	return &letter{}
}
