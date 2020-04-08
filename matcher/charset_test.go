package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"testing"
)


func TestNewCharset(t *testing.T) {

	useCases := []struct{
		description string
		charset string
		options []Option
		input []byte
		matched bool
	} {
		{
			description:"Byte match",
			charset:"abc",
			input:[]byte("b"),
			matched:true,
		},
		{
			description:"unicode match",
			charset:"日本語",
			input:[]byte("本"),
			matched:true,
		},
		{
			description:"unicode no match",
			charset:"日本語",
			input:[]byte("z"),
			matched:false,
		},

		{
			description:"case match",
			charset:"abc",
			input:[]byte("B"),
			matched:true,
			options:[]Option{
				&option.Case{Sensitive:false},
			},
		},
	}

	for _, useCase := range useCases {
		matcher := NewCharset(useCase.charset, useCase.options...)
		cursor := &parsly.Cursor{Input:useCase.input}
		matched := matcher.Match(cursor)
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}


