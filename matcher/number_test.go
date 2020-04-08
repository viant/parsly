package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewNumber(t *testing.T) {

	useCases := []struct{
		description string
		input []byte
		offset int
		matched int
	}{
		{
			description:"negative float",
			input:[]byte("-0.344  "),
			matched:6,
		},
		{
			description:"negative float pos",
			input:[]byte("  -0.344  "),
			offset:2,
			matched:6,
		},
		{
			description:"integer",
			input:[]byte("123 "),
			matched:3,
		},
		{
			description:"negative integer",
			input:[]byte("-123 "),
			matched:4,
		},
		{
			description:"invalid negative integer",
			input:[]byte("--123 "),
			matched:0,
		},
		{
			description:"float",
			input:[]byte("0.3 "),
			matched:3,
		},

		{
			description:"epx notation valid case 1",
			input:[]byte("0.3e-1 "),
			matched:6,
		},
		{
			description:"epx notation valid case 2",
			input:[]byte("10.3e+13 "),
			matched:8,
		},

		{
			description:"epx notation valid case 3",
			input:[]byte("-10.3e+13 "),
			matched:9,
		},
		{
			description:"epx notation valid case 4",
			input:[]byte("-10.3e1 "),
			matched:7,
		},
		{
			description:"invalid exp case 1",
			input:[]byte("-10.3ea "),
			matched:0,
		},
		{
			description:"invalid exp case 2",
			input:[]byte("-10.3e  "),
			matched:0,
		},
		{
			description:"invalid exp case 3",
			input:[]byte("-10.3e++  "),
			matched:0,
		},
		{
			description:"valid exp case 4",
			input:[]byte("-0.3e-11 日本語"),
			matched:8,
		},
		{
			description:"invalid exp case 5",
			input:[]byte("-0.3e"),
			matched:0,
		},
	}

	for _, useCase := range useCases {

		matcher := NewNumber()
		matched := matcher.Match(&parsly.Cursor{Input:useCase.input, Pos: useCase.offset})
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}

