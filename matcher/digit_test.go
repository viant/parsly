package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewDigit(t *testing.T) {

	useCases := []struct {
		description string
		options     []Option
		offset      int
		input       []byte
		matched     bool
	}{
		{
			description: "lower bound match",
			input:       []byte("0"),
			offset:      0,
			matched:     true,
		},
		{
			description: "upper bound match",
			input:       []byte("9"),
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
		matcher := NewDigit()
		matched := matcher.Match(&parsly.Cursor{Input:useCase.input, Pos:useCase.offset})
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}


