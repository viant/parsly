package matcher

import (
	"github.com/viant/parsly/lex"
	"unicode"
)

type number struct {
}

var digitMatcher = &digit{}
var decimalPoint = &char{value: '.'}
var minusMatcher = &char{value: '-'}
var signMatcher = &sign{}
var expMatcher = &charset{values: map[rune]bool{
	'e': true,
	'E': true,
}}

var expectDigitOrDecimalOrExp = []lex.RuneMatcher{digitMatcher, decimalPoint, expMatcher}
var expectDigit = []lex.RuneMatcher{digitMatcher}
var expectDigitOrDecimal = []lex.RuneMatcher{digitMatcher, decimalPoint}
var expectDigitOrExp = []lex.RuneMatcher{expMatcher, digitMatcher}

//Match matches a number
func (n *number) Match(input []byte, offset int) (matched int) {
	inputSize := len(input)
	runeValue := rune(input[offset])
	match := lex.MatchRune(runeValue, digitMatcher, minusMatcher)
	if match == nil {
		return 0
	}
	matched += 1
	isValidDigit := match == digitMatcher
	expect := expectDigit
	hasDecimal := false
	if isValidDigit {
		expect = expectDigitOrDecimalOrExp
	}
	for ; offset+matched < inputSize; {
		runeValue = rune(input[offset+matched])
		if match = lex.MatchRune(runeValue, expect...); match == nil {
			break
		}
		matched += 1
		hasMore := offset+matched < inputSize

		switch match {
		case digitMatcher:
			if ! isValidDigit && ! hasDecimal {
				expect = expectDigitOrDecimal
			}
			isValidDigit = true
			continue
		case expMatcher:
			if ! hasMore {
				return 0
			}
			runeValue = rune(input[offset+matched])
			isValidDigit = unicode.IsDigit(runeValue)
			if ! (signMatcher.MatchRune(runeValue) || isValidDigit) {
				return 0
			}
			matched += 1
			expect = expectDigit
			continue
		case decimalPoint:
			hasDecimal = true
			expect = expectDigitOrExp
			isValidDigit = false
		}
	}
	if ! isValidDigit {
		return 0
	}
	return matched
}

//NewNumber creates a number matcher
func NewNumber() lex.Matcher {
	return &number{

	}
}
