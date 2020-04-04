package matcher

import (
	"github.com/viant/parsly/lex"
)

type digit struct{}

func (d *digit) Match(input []byte, offset int) (matched int) {
	b := input[offset]
	if b < 48  || b > 57 {
		return 0
	}
	return 1
}

func (d *digit) MatchRune(runeValue rune) bool {
	if runeValue < 48 || runeValue > 57 {
		return false
	}
	return true
}

//NewDigit creates a digit matcher
func NewDigit() lex.Matcher {
	return &digit{}
}
