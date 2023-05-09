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
			description: "block with asterisks, and odd number of quotation marks (1)",
			begin:       '*',
			end:         '*',
			input:       []byte(`*"*`),
			pos:         0,
			matched:     3,
		},
		{
			description: "block with asterisks, and even number of quotation marks (2)",
			begin:       '*',
			end:         '*',
			input:       []byte(`*""*`),
			pos:         0,
			matched:     4,
		},
		{
			description: "block with asterisks, and odd number of quotation marks (3)",
			begin:       '*',
			end:         '*',
			input:       []byte(`*"""*`),
			pos:         0,
			matched:     5,
		},
		{
			description: "block with asterisks, and odd number of quotation marks (3) and additional spaces",
			begin:       '*',
			end:         '*',
			input:       []byte(`* " " " *`),
			pos:         0,
			matched:     9,
		},
		{
			description: "block with asterisks, and even number of quotation marks (4)",
			begin:       '*',
			end:         '*',
			input:       []byte(`*""""*`),
			pos:         0,
			matched:     6,
		},
		{
			description: "block with asterisks, and even number of quotation marks (4) and additional spaces",
			begin:       '*',
			end:         '*',
			input:       []byte(`* " " " " *`),
			pos:         0,
			matched:     11,
		},
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

		{
			description: "quteed",
			begin:       '\'',
			end:         '\'',
			escape:      '\\',
			input:       []byte(`'ew'rere`),
			matched:     4,
		},
	}

	for _, useCase := range useCases {
		matcher := NewBlock(useCase.begin, useCase.end, useCase.escape)
		cursor := &parsly.Cursor{Input: useCase.input, Pos: useCase.pos}
		matched := matcher.Match(cursor)
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}
