package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewDigits(t *testing.T) {

	useCases := []struct {
		description string
		options     []Option
		offset      int
		input       []byte
		matched     bool
	}{
		{
			description: "lower bound match",
			input:       []byte("03"),
			offset:      0,
			matched:     true,
		},
		{
			description: "upper bound match",
			input:       []byte("93"),
			offset:      0,
			matched:     true,
		},
		{
			description: "no match",
			input:       []byte("z"),
			offset:      0,
			matched:     false,
		},
	}
	for _, useCase := range useCases {
		matcher := NewDigits()
		matched := matcher.Match(&parsly.Cursor{Input: useCase.input, Pos: useCase.offset})
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
		if !useCase.matched {
			continue
		}
		assert.EqualValues(t, len(useCase.input), matched, useCase.description)
	}

}
