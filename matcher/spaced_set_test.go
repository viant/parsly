package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewSpacedSet(t *testing.T) {
	useCases := []struct {
		description string
		fragments   []string
		options     []Option
		input       []byte
		matched     bool
		offset      int
	}{

		{
			description: "seq1 seq3 match",
			input:       []byte("seq1 seq3"),
			fragments: []string{
				"seq1 seq2",
				"seq1 seq3",
				"seq1",
			},
			matched: true,
		},

		{
			description: "seq1 seq3 match",
			input:       []byte("seq1 seq3"),
			fragments: []string{
				"seq1 seq2",
				"seq1 seq3",
			},
			matched: true,
		},
		{
			description: "seq1 seq4 not match",
			input:       []byte("seq1 seq4"),
			fragments: []string{
				"seq1 seq2",
				"seq1 seq3",
			},
			matched: false,
		},
		{
			description: "other match",
			input:       []byte("other"),
			fragments: []string{
				"seq1 seq2",
				"seq1 seq3",
				"other",
			},
			matched: true,
		},

		{
			description: "seq1 match",
			input:       []byte("seq1 ddddddddddd"),
			fragments: []string{
				"seq1 seq2",
				"seq1 seq3",
				"seq1",
			},
			matched: true,
		},
	}

	for _, useCase := range useCases {
		matcher := NewSpacedSet(useCase.fragments, useCase.options...)
		matched := matcher.Match(parsly.NewCursor("", useCase.input, useCase.offset))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
