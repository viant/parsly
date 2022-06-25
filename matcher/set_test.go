package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewSet(t *testing.T) {
	useCases := []struct {
		description string
		fragments   [][]byte
		options     []Option
		input       []byte
		matched     bool
	}{
		{
			description: "Set match",
			input:       []byte("asc"),
			fragments: [][]byte{
				[]byte("asc"),
				[]byte("desc"),
			},
			matched: true,
		},
		{
			description: "Set match",
			input:       []byte("desc"),
			fragments: [][]byte{
				[]byte("asc"),
				[]byte("desc"),
			},
			matched: true,
		},
		{
			description: "Set match",
			input:       []byte("desc3"),
			fragments: [][]byte{
				[]byte("asc"),
				[]byte("desc"),
			},
			matched: true,
		},
		{
			description: "Set match",
			input:       []byte("zzzzz"),
			fragments: [][]byte{
				[]byte("asc"),
				[]byte("desc"),
			},
			matched: false,
		},
		{
			description: "Set match",
			input:       []byte("x"),
			fragments: [][]byte{
				[]byte("asc"),
				[]byte("desc"),
			},
			matched: false,
		},
	}

	for _, useCase := range useCases {
		matcher := NewSpaceFragment(useCase.fragments, useCase.options...)
		matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
