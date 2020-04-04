package matcher

import (
	"bytes"
	"github.com/viant/parsly/lex"
	"github.com/viant/parsly/matcher/option"
)

type fragment struct {
	value []byte
	size  int
}

func (d *fragment) Match(input []byte, offset int) int {
	if offset+d.size < len(input) {
		if bytes.EqualFold(input[offset:offset+d.size], d.value) {
			return d.size
		}
	}
	return 0
}

type caseSensitiveFragment struct {
	value []byte
	size  int
}

func (d *caseSensitiveFragment) Match(input []byte, offset int)  int {
	if offset+d.size < len(input) {
		if bytes.Equal(input[offset:offset+d.size], d.value) {
			return d.size
		}
	}
	return 0
}

//NewFragments creates fragment matcher
func NewFragment(value string, options ...lex.Option) lex.Matcher {
	caseOpt := &option.Case{}
	if lex.AssignOption(options, &caseOpt) && ! caseOpt.Sensitive {
		return &fragment{
			value: []byte(value),
			size:  len(value),
		}
	}
	return &caseSensitiveFragment{
		value: []byte(value),
		size:  len(value),
	}
}
