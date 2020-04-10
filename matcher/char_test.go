package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewChar(t *testing.T) {

	useCases := []struct {
		description string
		char        rune
		options     []Option
		pos         int
		input       []byte
		matched     bool
	}{
		{
			description: "match",
			char:        'b',
			input:       []byte("abc"),
			pos:         1,
			matched:     true,
		},
		{
			description: "no match",
			char:        'b',
			pos:         0,
			input:       []byte("abc"),
			matched:     false,
		},
		{
			description: "unicode input",
			char:        'b',
			pos:         1,
			input:       []byte("ab日本語"),
			matched:     true,
		},

		{
			description: "unicode match",
			char:        '日',
			pos:         2,
			input:       []byte("ab日本語"),
			matched:     true,
		},
	}

	for _, useCase := range useCases {
		matcher := NewChar(useCase.char)
		cursor := &parsly.Cursor{Input: useCase.input, Pos: useCase.pos}
		matched := matcher.Match(cursor)
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
