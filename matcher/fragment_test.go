package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"testing"
)

func TestNewFragment(t *testing.T) {

	useCases := []struct {
		description string
		fragments   string
		options     []Option
		input       []byte
		matched     bool
	}{
		{
			description: "FragmentsFold match",
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
			options: []Option{
				&option.Case{Sensitive: false},
			},
		},
	}

	for _, useCase := range useCases {
		matcher := NewFragment(useCase.fragments, useCase.options...)
		matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
