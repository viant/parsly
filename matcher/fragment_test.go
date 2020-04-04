package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly/lex"
	"github.com/viant/parsly/matcher/option"
	"testing"
)

func TestNewFragment(t *testing.T) {

	useCases := []struct{
		description string
		fragments   string
		options     []lex.Option
		input       []byte
		matched     bool
	} {
		{
			description: "fragments match",
			fragments:   "abc",
			input:       []byte("abc test"),
			matched:     true,
		},
		{
			description: "unicode match",
			fragments:   "日本語",
			input:       []byte("日本語 test"),
			matched:     true,
		},
		{
			description: "unicode no match",
			fragments:   "日本語",
			input:       []byte("日本 test"),
			matched:     false,
		},

		{
			description: "case match",
			fragments:   "abc",
			input:       []byte("ABc test"),
			matched:     true,
			options:[]lex.Option{
				&option.Case{Sensitive:false},
			},
		},
	}

	for _, useCase := range useCases {
		matcher := NewFragment(useCase.fragments, useCase.options...)
		matched := matcher.Match(useCase.input, 0)
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}



}
