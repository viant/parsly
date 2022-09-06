package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"testing"
)

func TestNewSpacedFragment(t *testing.T) {
	useCases := []struct {
		description string
		fragments   string
		options     []Option
		input       []byte
		offset      int
		expect      int
	}{

		{
			description: "match case 1",
			input:       []byte("  order \nby rer"),
			fragments:   " order by",
			offset:      1,
			expect:      10,
		},

		{
			description: "no match case 3",
			input:       []byte(" datainto zxx"),
			offset:      1,
			fragments:   "data into",
			options:     []Option{&option.Case{Sensitive: false}},
			expect:      0,
		},

		{
			description: "match case 2",
			input:       []byte("order \tby rer"),
			fragments:   "order by",
			expect:      9,
		},
		{
			description: "match case 3",
			input:       []byte("order\tby s"),
			fragments:   "order by",
			expect:      8,
		},
		{
			description: "no match case 1",
			input:       []byte("order\tbz s"),
			fragments:   "order by",
			expect:      0,
		},
		{
			description: "no match case 2",
			input:       []byte(" datainto zxx"),
			offset:      1,
			fragments:   "data into",
			options:     []Option{&option.Case{Sensitive: false}},
			expect:      0,
		},
		{
			description: "no match case 3",
			input:       []byte(" datainto zxx"),
			offset:      1,
			fragments:   "data into",
			options:     []Option{&option.Case{Sensitive: true}},
			expect:      0,
		},
		{
			description: "no match case 4",
			input:       []byte("order\tbz s"),
			options:     []Option{&option.Case{Sensitive: true}},
			fragments:   "order by",
			expect:      0,
		},
	}

	for _, useCase := range useCases {
		matcher := NewSpacedFragment(useCase.fragments, useCase.options...)
		cursor := parsly.NewCursor("", useCase.input, useCase.offset)
		cursor.Pos = useCase.offset
		matched := matcher.Match(cursor)
		assert.Equal(t, useCase.expect, matched, useCase.description)
	}

}
