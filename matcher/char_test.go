package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly/lex"
	"testing"
)

func TestNewChar(t *testing.T) {

	useCases := []struct {
		description string
		char     rune
		options     []lex.Option
		offset int
		input       []byte
		matched     bool
	}{
		{
			description:"match",
			char:'b',
			input:[]byte("abc"),
			offset:1,
			matched:true,
		},
		{
			description:"no match",
			char:'b',
			offset:0,
			input:[]byte("abc"),
			matched:false,
		},
		{
			description:"unicode input",
			char:'b',
			offset:1,
			input:[]byte("ab日本語"),
			matched:true,
		},

		{
			description:"unicode match",
			char:'日',
			offset:2,
			input:[]byte("ab日本語"),
			matched:true,
		},


	}

	for _, useCase := range useCases {
		matcher := NewChar(useCase.char)
		matched := matcher.Match(useCase.input, useCase.offset)
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}