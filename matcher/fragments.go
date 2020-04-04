package matcher

import (
	"bytes"
	"github.com/viant/parsly/lex"
	"github.com/viant/parsly/matcher/option"
)

type fragments struct {
	values [][]byte
}

func (d *fragments) Match(input []byte, offset int) (matched int) {
	for _, candidate := range d.values {
		if offset+len(candidate) < len(input) {
			if bytes.EqualFold(input[offset:offset+len(candidate)], candidate) {
				return len(candidate)
			}
		}
	}
	return matched
}


type caseSensitiveFragments struct {
	values [][]byte
}

func (d *caseSensitiveFragments) Match(input []byte, offset int) (matched int) {
	for _, candidate := range d.values {
		if offset+len(candidate) < len(input) {
			if bytes.Equal(input[offset:offset+len(candidate)], candidate) {
				return len(candidate)
			}
		}
	}
	return matched
}

//NewFragments creates fragment matcher
func NewFragments(values []string, options ...lex.Option) lex.Matcher {
	var vals = make([][]byte, 0)
	for _, value := range values {
		vals = append(vals, []byte(value))
	}
	caseOpt := &option.Case{}
	if lex.AssignOption(options, &caseOpt) && ! caseOpt.Sensitive {
		return &fragments{values: vals}
	}
	return &caseSensitiveFragments{values: vals}
}
