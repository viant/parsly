package matcher

import (
	"bytes"
	"github.com/viant/parsly"
)

type FragmentsFold struct {
	values [][]byte
}

func (d *FragmentsFold) Match(cursor *parsly.Cursor) (matched int) {
	size := cursor.InputSize
	offset := cursor.Pos
	input := cursor.Input
	for _, candidate := range d.values {
		matchEnd := offset+len(candidate)
		if matchEnd < size {
			if bytes.EqualFold(input[offset:matchEnd], candidate) {
				return len(candidate)
			}
		}
	}
	return matched
}

//NewFragmentsFold returns fragments folds
func NewFragmentsFold(values ...[]byte) *FragmentsFold {
	return &FragmentsFold{values:values}
}



type Fragments  struct {
	values [][]byte
}

func (d *Fragments) Match(cursor *parsly.Cursor) (matched int) {
	size := cursor.InputSize
	offset := cursor.Pos
	input := cursor.Input
	for _, candidate := range d.values {
		matchEnd := offset+len(candidate)
		if matchEnd < size {
			if bytes.Equal(input[offset:matchEnd], candidate) {
				return len(candidate)
			}
		}
	}
	return matched
}



//NewFragments returns fragments folds
func NewFragments(values ...[]byte) *Fragments {
	return &Fragments{values:values}
}



