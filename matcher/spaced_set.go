package matcher

import (
	"bytes"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"strings"
)

type SpaceSetFold struct {
	set [][][]byte
}

func (d *SpaceSetFold) Match(cursor *parsly.Cursor) int {
	pos := cursor.Pos
	inputLen := len(cursor.Input)
	matched := 0
next:
	for _, items := range d.set {
		pos = cursor.Pos
		matched = 0
		for i, value := range items {
			if !MatchFold(value, cursor.Input, 0, pos) {
				continue next
			}
			pos += len(value)
			matched += len(value)
			isLast := i == len(items)-1
			if isLast {
				return matched
			}
			for j := pos; j < inputLen-1; j++ {
				if !IsWhiteSpace(cursor.Input[pos]) {
					if j > 0 {
						break
					}
					pos = 0
					matched = 0
					break next
				}
				matched++
				pos++
			}
		}
	}
	return 0
}

//SpacedSet represent space fragment
type SpacedSet struct {
	set [][][]byte
}

func (d *SpacedSet) Match(cursor *parsly.Cursor) int {
	pos := cursor.Pos
	inputLen := len(cursor.Input)
	matched := 0
next:
	for _, items := range d.set {
		pos = cursor.Pos
		matched = 0
		for i, value := range items {
			if !bytes.Equal(value, cursor.Input[pos:pos+len(value)]) {
				continue next
			}
			pos += len(value)
			matched += len(value)
			isLast := i == len(items)-1
			if isLast {
				return matched
			}
			for j := pos; j < inputLen-1; j++ {
				if !IsWhiteSpace(cursor.Input[pos]) {
					if j > 0 {

						break
					}
					pos = 0
					matched = 0
					break next
				}
				matched++
				pos++
			}
		}
	}
	return 0
}

//NewSpacedSet create a spaced set
func NewSpacedSet(texts []string, options ...Option) parsly.Matcher {
	caseOpt := &option.Case{}
	var set = make([][][]byte, 0)
	for _, text := range texts {
		values := strings.Split(text, " ")
		size := len(values) - 1
		var spaced = make([][]byte, 0)
		for i, v := range values {
			size += len(v)
			spaced = append(spaced, []byte(values[i]))
		}
		set = append(set, spaced)
	}
	if AssignOption(options, &caseOpt) && !caseOpt.Sensitive {
		return &SpaceSetFold{
			set: set,
		}
	}
	return &SpacedSet{
		set: set,
	}
}
