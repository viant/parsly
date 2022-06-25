package matcher

import (
	"bytes"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
)

type SetFold struct {
	value [][]byte
}

func (d *SetFold) Match(cursor *parsly.Cursor) int {
	for _, value := range d.value {
		matchEnd := cursor.Pos + len(value)
		if matchEnd < cursor.InputSize {
			if MatchFold(value, cursor.Input, 0, cursor.Pos) {
				return len(value)
			}
		}
	}
	return 0
}

type Set struct {
	value [][]byte
	size  int
}

func (d *Set) Match(cursor *parsly.Cursor) int {
	for _, value := range d.value {
		matchEnd := cursor.Pos + len(value)
		if matchEnd < cursor.InputSize {
			if bytes.Equal(value, cursor.Input[cursor.Pos:matchEnd]) {
				return len(value)
			}
		}
	}
	return 0
}

//NewSets creates SetFold matcher
func NewSet(value [][]byte, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	if AssignOption(options, &caseOpt) && !caseOpt.Sensitive {
		return &SetFold{
			value: value,
		}
	}
	return &Set{
		value: value,
	}
}
