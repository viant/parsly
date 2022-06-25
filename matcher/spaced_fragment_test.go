package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewSpacedFragment(t *testing.T) {
	useCases := []struct {
		description string
		fragments   [][]byte
		options     []Option
		input       []byte
		matched     bool
	}{
		{
			description: "FragmentsFold match end",
			input:       []byte("order\tby"),
			fragments: [][]byte{
				[]byte("order"),
				[]byte("by"),
			},
			matched: true,
		},
		{
			description: "FragmentsFold match",
			input:       []byte("order\tby s"),
			fragments: [][]byte{
				[]byte("order"),
				[]byte("by"),
			},
			matched: true,
		},
		{
			description: "FragmentsFold no match",
			input:       []byte("order\tbz s"),
			fragments: [][]byte{
				[]byte("order"),
				[]byte("by"),
			},
			matched: true,
		},
	}

	for _, useCase := range useCases {
		matcher := NewSpaceFragment(useCase.fragments, useCase.options...)
		matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
