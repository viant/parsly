package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewSet(t *testing.T) {
	useCases := []struct {
		description string
		fragments   []string
		options     []Option
		input       []byte
		matched     bool
	}{
		{
			description: "Set match 1 ",
			input:       []byte("asc"),
			fragments: []string{
				"asc",
				"desc",
			},
			matched: true,
		},
		{
			description: "Set match 2 ",
			input:       []byte("desc"),
			fragments: []string{
				"asc",
				"desc",
			},
			matched: true,
		},
		{
			description: "Set match 3",
			input:       []byte("desc3"),
			fragments: []string{
				"asc",
				"desc",
			},
			matched: true,
		},
		{
			description: "Set no match 1 ",
			input:       []byte("zzzzz"),
			fragments: []string{
				"asc",
				"desc",
			},
			matched: false,
		},
		{
			description: "Set not match 2",
			input:       []byte("x"),
			fragments: []string{
				"asc",
				"desc",
			},
			matched: false,
		},
	}

	for _, useCase := range useCases {
		matcher := NewSet(useCase.fragments, useCase.options...)
		matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
