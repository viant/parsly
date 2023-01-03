package splitter

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly/matcher"
	"github.com/viant/parsly/matcher/option"
	"testing"
)

func TestSplit(t *testing.T) {
	var testCases = []struct {
		description string
		delimiters  []string
		options     []matcher.Option
		text        string
		expect      []string
	}{
		{
			description: "case sensitive split",
			delimiters:  []string{"W1", "w2"},
			text:        "abw1cdW1dddw2eee",

			expect: []string{"abw1cd", "ddd", "eee"},
		},
		{
			description: "case insensitive split",
			delimiters:  []string{"W1", "w2"},
			text:        "abw1cdW1dddw2eee",
			options:     []matcher.Option{option.NewCase(false)},
			expect:      []string{"ab", "cd", "ddd", "eee"},
		},
	}

	for _, testCase := range testCases {
		splitter := New(testCase.delimiters, testCase.options...)
		actual := splitter.Split(testCase.text)
		assert.EqualValues(t, testCase.expect, actual, testCase.description)
	}
}
