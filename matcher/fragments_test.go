package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewFragments(t *testing.T) {

	useCases := []struct{
		description     string
		fragments       string
		caseInsensitive bool
		input           []byte
		matched         bool
	} {
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
			description:     "case match",
			fragments:       "abc",
			input:           []byte("ABc test"),
			matched:         true,
			caseInsensitive: true,
		},
	}

	for _, useCase := range useCases {


		if useCase.caseInsensitive {
			matcher := NewFragmentsFold([]byte(useCase.fragments))
			matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
			assert.Equal(t, useCase.matched, matched > 0, useCase.description)
			continue
		}

		matcher := NewFragments([]byte(useCase.fragments))
		matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)

	}



}
