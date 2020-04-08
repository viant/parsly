package matcher

import (
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"unicode"
)


//NewCharset creates a bytes or runes matcher
func NewCharset(set string, options ...Option) parsly.Matcher {
	valuesMap := make(map[rune]bool)

	caseOpt := &option.Case{}
	if ! AssignOption(options, &caseOpt) {
		caseOpt = nil
	}
	useRunes := false
	for _, r := range set {
		if ! useRunes && !isByte(r) {
			useRunes = true
		}
		if caseOpt != nil && ! caseOpt.Sensitive {
			valuesMap[unicode.ToLower(r)] = true
			valuesMap[unicode.ToUpper(r)] = true
			continue
		}
		valuesMap[r] = true
	}
	if useRunes {
		values := make([]rune, 0)
		for k := range valuesMap {
			values = append(values, k)
		}
		return NewRunes(values)
	}

	var values = make([]byte, 0)
	for k := range valuesMap {
		values = append(values, byte(k))
	}
	return NewBytes(values)
}
