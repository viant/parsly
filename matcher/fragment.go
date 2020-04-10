package matcher

import (
	"bytes"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
)

type FragmentFold struct {
	value []byte
	size  int
}

func (d *FragmentFold) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	if matchEnd < cursor.InputSize {
		if bytes.EqualFold(cursor.Input[cursor.Pos:matchEnd], d.value) {
			return d.size
		}
	}
	return 0
}

type Fragment struct {
	value []byte
	size  int
}

func (d *Fragment) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	if matchEnd < cursor.InputSize {
		if bytes.Equal(cursor.Input[cursor.Pos:matchEnd], d.value) {
			return d.size
		}
	}
	return 0
}

//NewFragments creates FragmentFold matcher
func NewFragment(value string, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	if AssignOption(options, &caseOpt) && !caseOpt.Sensitive {
		return &FragmentFold{
			value: []byte(value),
			size:  len(value),
		}
	}
	return &Fragment{
		value: []byte(value),
		size:  len(value),
	}
}
