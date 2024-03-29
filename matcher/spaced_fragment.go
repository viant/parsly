package matcher

import (
	"bytes"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"strings"
)

type SpaceFragmentFold struct {
	values [][]byte
	size   int
}

func (d *SpaceFragmentFold) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	originalPos := cursor.Pos
	pos := cursor.Pos
	inputLen := len(cursor.Input)
	if matchEnd > inputLen {
		return 0
	}

	for i, value := range d.values {
		if !MatchFold(value, cursor.Input, 0, pos) {
			return 0
		}
		pos += len(value)

		if i == len(d.values)-1 { //last seq matched
			break
		}

		if i != len(d.values)-1 {
			spaceCounter := 0
			for j := pos; j < inputLen-1; j++ {
				if !IsWhiteSpace(cursor.Input[j]) {
					if spaceCounter > 0 {
						break
					}
					return 0
				}
				spaceCounter++
				pos++
			}
		}
	}
	return pos - originalPos
}

//SpacedFragment represent space fragment
type SpacedFragment struct {
	values [][]byte
	size   int
}

func (d *SpacedFragment) Match(cursor *parsly.Cursor) int {
	matchEnd := cursor.Pos + d.size
	pos := cursor.Pos
	originalPos := cursor.Pos

	inputLen := len(cursor.Input)
	if matchEnd > inputLen {
		return 0
	}

	for i, value := range d.values {
		if !bytes.Equal(value, cursor.Input[pos:pos+len(value)]) {
			return 0
		}
		pos += len(value)

		if i == len(d.values)-1 { //last seq matched
			break
		}
		spaceCounter := 0

		for j := pos; j < inputLen-1; j++ {
			if !IsWhiteSpace(cursor.Input[j]) {
				if spaceCounter > 0 {
					break
				}
				return 0
			}
			pos++
			spaceCounter++
		}

	}
	return pos - originalPos
}

//NewSpaceFragments creates SpaceFragmentFold matcher
func NewSpacedFragment(text string, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	values := strings.Split(text, " ")
	size := len(values) - 1
	var fragments = make([][]byte, 0)
	for i, v := range values {
		size += len(v)
		fragments = append(fragments, []byte(values[i]))
	}
	if AssignOption(options, &caseOpt) && !caseOpt.Sensitive {
		return &SpaceFragmentFold{
			values: fragments,
			size:   size,
		}
	}
	return &SpacedFragment{
		values: fragments,
		size:   size,
	}
}
