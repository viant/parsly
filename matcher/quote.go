package matcher

import (
	"github.com/viant/parsly/lex"
	"unicode/utf8"
)

type utfQuote struct {
	escape rune
	value  rune
}

//Match matches quoted characters
func (m *utfQuote) Match(input []byte, offset int) int {
	var matched = 0
	inputSize:=len(input)
	runeValue, width := utf8.DecodeRune(input[offset:])
	if runeValue != m.value {
		return 0
	}
	matched += width
	for i := offset  + matched; i < inputSize ; i++ {
		runeValue, width = utf8.DecodeRune(input[offset+matched:])
		matched += width
		hasMore := offset+matched < inputSize
		if runeValue == m.escape {
			if ! hasMore {
				return 0
			}
			runeValue, width = utf8.DecodeRune(input[offset+matched:])
			matched += width
			continue
		}
		if runeValue == m.value {
			return matched
		}
	}
	return 0
}

type quote struct {
	escape byte
	value  byte
}

//Match matches quoted characters
func (m *quote) Match(input []byte, offset int) int {
	var matched = 0
	inputSize:=len(input)
	value := input[offset]
	if value != m.value {
		return 0
	}
	matched++
	for i := offset + matched; i < inputSize ; i++ {
		value = input[i]
		matched++
		hasMore := i + 1 < inputSize
		if value == m.escape {
			if ! hasMore {
				return 0
			}
			i++
			matched++
			continue
		}
		if value == m.value {
			return matched
		}
	}
	return 0
}


//NewQuote creates a new utfQuote matcher
func NewQuote(val, escape rune) lex.Matcher {
	if isByte(val) && isByte(escape) {
		return &quote{
			escape: byte(escape),
			value:  byte(val),
		}
	}
	return &utfQuote{
		value:  val,
		escape: escape,
	}
}
