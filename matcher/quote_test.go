package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewQuote(t *testing.T) {
	useCases := []struct {
		description string
		quote       rune
		escape      rune
		input       []byte
		matched     int
	}{

		{
			description: "RuneQuote with unicode Byte match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'日本語 test' abc"),
			matched:     16,
		},
		{
			description: "RuneQuote with unicode Byte match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'this is test' abc"),
			matched:     14,
		},
		{
			description: "RuneQuote with unicode Byte match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'t \\'is test' abc"),
			matched:     13,
		},

		{
			description: "RuneQuote with quote and escape the same",
			quote:       '\'',
			escape:      '\'',
			input:       []byte("'abc''r'3434"),
			matched:     8,
		},
		{
			description: "RuneQuote with quote and escape the same at the end",
			quote:       '\'',
			escape:      '\'',
			input:       []byte("'abc''r'"),
			matched:     8,
		},
	}

	for _, useCase := range useCases {
		matcher := NewRuneQuote(useCase.quote, useCase.escape)
		matched := matcher.Match(&parsly.Cursor{Input: useCase.input})
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}

func TestNewByteQuote(t *testing.T) {
	useCases := []struct {
		description string
		quote       byte
		escape      byte
		input       []byte
		matched     int
	}{

		{
			description: "RuneQuote with unicode Byte match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'this is test' abc"),
			matched:     14,
		},
		{
			description: "RuneQuote with unicode Byte match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'t \\'is test' abc"),
			matched:     13,
		},

		{
			description: "RuneQuote with quote and escape the same",
			quote:       '\'',
			escape:      '\'',
			input:       []byte("'abc''r'3434"),
			matched:     8,
		},
	}

	for _, useCase := range useCases {
		matcher := NewByteQuote(useCase.quote, useCase.escape)
		matched := matcher.Match(&parsly.Cursor{Input: useCase.input})
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}
