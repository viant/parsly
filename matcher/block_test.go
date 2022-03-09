package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewBlock(t *testing.T) {

	useCases := []struct {
		description string
		begin       byte
		end         byte
		escape      byte
		pos         int
		input       []byte
		matched     int
	}{
		{
			description: "match",
			begin:       '{',
			end:         '}',
			escape:      '\\',
			input:       []byte("x{test{2}xe}vc"),
			pos:         1,
			matched:     11,
		},
		{
			description: "match",
			begin:       '{',
			end:         '}',
			escape:      '\\',
			input:       []byte(`x{test{"{2"}xe}vc`),
			pos:         1,
			matched:     14,
		},
		{
			description: "begin and end the same",
			begin:       ';',
			end:         ';',
			escape:      '\\',
			input:       []byte(`;123;`),
			matched:     5,
		},
	}

	for _, useCase := range useCases {
		matcher := NewBlock(useCase.begin, useCase.end, useCase.escape)
		cursor := &parsly.Cursor{Input: useCase.input, Pos: useCase.pos}
		matched := matcher.Match(cursor)
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}
