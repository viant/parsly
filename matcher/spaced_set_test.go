package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"testing"
)

func TestNewSpacedSet(t *testing.T) {
	useCases := []struct {
		description string
		fragments   []string
		options     []Option
		input       []byte
		matched     bool
		pos         int
	}{

		//{
		//	description: "seq1 seq3 match",
		//	input:       []byte("seq1 seq3"),
		//	fragments: []string{
		//		"seq1 seq2",
		//		"seq1 seq3",
		//		"seq1",
		//	},
		//	matched: true,
		//},
		//
		//{
		//	description: "seq1 seq3 match",
		//	input:       []byte("seq1 seq3"),
		//	fragments: []string{
		//		"seq1 seq2",
		//		"seq1 seq3",
		//	},
		//	matched: true,
		//},
		//{
		//	description: "seq1 seq4 not match",
		//	input:       []byte("seq1 seq4"),
		//	fragments: []string{
		//		"seq1 seq2",
		//		"seq1 seq3",
		//	},
		//	matched: false,
		//},
		//{
		//	description: "other match",
		//	input:       []byte("other"),
		//	fragments: []string{
		//		"seq1 seq2",
		//		"seq1 seq3",
		//		"other",
		//	},
		//	matched: true,
		//},
		{
			description: "seq1 match",
			input:       []byte("x seq1 ddddddddddd"),
			pos:         2,
			options: []Option{
				&option.Case{},
			},
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
		cur := parsly.NewCursor("", useCase.input, 0)
		cur.Pos = useCase.pos
		matched := matcher.Match(cur)
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
