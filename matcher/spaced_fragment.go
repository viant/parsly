package matcher

import (
	"bytes"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
)

type SpaceFragmentFold struct {
	values [][]byte
	size   int
}

func (d *SpaceFragmentFold) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	pos := cursor.Pos
	inputLen := len(cursor.Input)
	if matchEnd > inputLen {
		return 0
	}
outer:
	for i, value := range d.values {
		if !MatchFold(value, cursor.Input, 0, pos) {
			return 0
		}
		pos += len(value)
		if i != len(d.values)-1 {
			for j := pos; j < inputLen-1; j++ {
				if !IsWhiteSpace(cursor.Input[pos]) {
					if j > 0 {
						break outer
					}
					return 0
				}
				pos++
			}
		}
	}
	return pos
}

//SpacedFragment represent space fragment
type SpacedFragment struct {
	values [][]byte
	size   int
}

func (d *SpacedFragment) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	pos := cursor.Pos
	inputLen := len(cursor.Input)
	if matchEnd > inputLen {
		return 0
	}
outer:
	for i, value := range d.values {
		if !bytes.Equal(value, cursor.Input[pos:pos+len(value)]) {
			return 0
		}
		pos += len(value)

		if i == len(d.values)-1 {
			break
		}

		for j := pos; j < inputLen-1; j++ {
			if !IsWhiteSpace(cursor.Input[pos]) {
				if j > 0 {
					break outer
				}
				return 0
			}
			pos++
		}

	}
	return pos
}

//NewSpaceFragments creates SpaceFragmentFold matcher
func NewSpaceFragment(value [][]byte, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	size := len(value) - 1
	for _, v := range value {
		size += len(v)
	}
	if AssignOption(options, &caseOpt) && !caseOpt.Sensitive {
		return &SpaceFragmentFold{
			values: value,
			size:   size,
		}
	}
	return &SpacedFragment{
		values: value,
		size:   size,
	}
}
