package matcher

import (
	"github.com/stretchr/testify/assert"
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
			description: "utfQuote with unicode utfChar match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'日本語 test' abc"),
			matched:     16,
		},
		{
			description: "utfQuote with unicode utfChar match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'this is test' abc"),
			matched:     14,
		},
		{
			description: "utfQuote with unicode utfChar match",
			quote:       '\'',
			escape:      '\\',
			input:       []byte("'t \\'is test' abc"),
			matched:     13,
		},
	}

	for _, useCase := range useCases {
		matcher := NewQuote(useCase.quote, useCase.escape)
		matched := matcher.Match(useCase.input, 0)
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}
