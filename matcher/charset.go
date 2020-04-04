package matcher

import (
	"github.com/viant/parsly/lex"
	"github.com/viant/parsly/matcher/option"
	"unicode"
	"unicode/utf8"
)

//utfCharset represents a matcher, that matches any of runes.
type utfCharset struct {
	values map[rune]bool
}

//Match matches any characters defined in Chars in the input, returns 1 if character has been matched
func (m *utfCharset) Match(input []byte, offset int) int {
	runeValue, width := utf8.DecodeRune(input[offset:])
	if m.MatchRune(runeValue) {
		return width
	}
	return 0
}

//MatchRune matches rune
func (m *utfCharset) MatchRune(runeValue rune) bool {
	return m.values[runeValue]
}

//utfCharset represents a matcher, that matches any of runes.
type charset struct {
	values map[rune]bool
}

//Match matches any characters defined in Chars in the input, returns 1 if character has been matched
func (m *charset) Match(input []byte, offset int) int {
	value := rune(input[offset])
	if m.MatchRune(value) {
		return 1
	}
	return 0
}

//MatchRune matches rune
func (m *charset) MatchRune(runeValue rune) bool {
	return m.values[runeValue]
}

//NewCharset creates a utfCharset matcher
func NewCharset(set string, options ...lex.Option) lex.Matcher {
	values := make(map[rune]bool)

	caseOpt := &option.Case{}
	if ! lex.AssignOption(options, &caseOpt) {
		caseOpt = nil
	}
	useUTF := false
	for _, r := range set {
		if ! useUTF && !isByte(r) {
			useUTF = true
		}
		if caseOpt != nil && ! caseOpt.Sensitive {
			values[unicode.ToLower(r)] = true
			values[unicode.ToUpper(r)] = true
			continue
		}
		values[r] = true
	}
	if useUTF {
		return &utfCharset{values: values}
	}
	return &charset{values: values}
}
