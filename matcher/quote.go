package matcher

import (
	"github.com/viant/parsly"
	"unicode/utf8"
)

type RuneQuote struct {
	escape rune
	value  rune
}

//TokenMatch matches quoted characters
func (m *RuneQuote) Match(cursor *parsly.Cursor) int {
	var matched = 0
	input := cursor.Input
	inputSize := len(input)

	pos := cursor.Pos
	runeValue, width := utf8.DecodeRune(input[pos:])
	if runeValue != m.value {
		return 0
	}
	matched += width
	escapeQuoteTheSame := m.escape == m.value

	for i := pos + matched; i < inputSize; i++ {
		runeValue, width = utf8.DecodeRune(input[pos+matched:])
		matched += width
		hasMore := pos+matched < inputSize

		if escapeQuoteTheSame {
			if runeValue == m.escape {
				runeValue, width = utf8.DecodeRune(input[pos+matched:])
				if hasMore && runeValue == m.value {
					matched += width
					continue
				}
				return matched
			}
		} else {
			if runeValue == m.escape {
				if !hasMore {
					return 0
				}
				runeValue, width = utf8.DecodeRune(input[pos+matched:])
				matched += width
				continue
			}
		}
		if runeValue == m.value {
			return matched
		}
	}
	return 0
}

func NewRuneQuote(value, escape rune) *RuneQuote {
	return &RuneQuote{
		escape: escape,
		value:  value,
	}
}

type ByteQuote struct {
	escape byte
	value  byte
}

//TokenMatch matches quoted characters
func (m *ByteQuote) Match(cursor *parsly.Cursor) int {
	var matched = 0
	input := cursor.Input
	inputSize := len(cursor.Input)
	pos := cursor.Pos
	value := input[pos]
	if value != m.value {
		return 0
	}
	matched++
	escapeQuoteTheSame := m.escape == m.value

	for i := pos + matched; i < inputSize; i++ {
		value = input[i]
		matched++
		hasMore := i+1 < inputSize
		if escapeQuoteTheSame {
			if value == m.escape {
				if hasMore && input[i+1] == m.value {
					i++
					matched++
					continue
				}
				return matched
			}
		} else {
			if value == m.escape {
				if !hasMore {
					return 0
				}
				i++
				matched++
				continue
			}
		}

		if value == m.value {
			return matched
		}
	}
	return 0
}

//NewByteQuote creates a byte quote
func NewByteQuote(value, escape byte) *ByteQuote {
	return &ByteQuote{
		escape: escape,
		value:  value,
	}
}

//NewQuote creates a new RuneQuote matcher
func NewQuote(val, escape rune) parsly.Matcher {
	if isByte(val) && isByte(escape) {
		return NewByteQuote(byte(val), byte(escape))
	}
	return NewRuneQuote(val, escape)
}
