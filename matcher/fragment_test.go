package matcher

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher/option"
	"strings"
	"testing"
)

func genInput(length int) []byte {
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = (byte)(i)
	}
	return result
}

func TestMatchFoldFragment(t *testing.T) {
	for i := 0; i < 1000; i++ { // all matches
		s1 := strings.Repeat("a", i+1)
		s2 := strings.Repeat("A", i+1)
		assert.True(t, MatchFold([]byte(s1), []byte(s2), 0, 0), fmt.Sprintf("fialed for i: %v", i+1))
	}

	assert.True(t, MatchFold([]byte("abcde"), []byte("zxabcdefg"), 0, 2))

	for i := 0; i < 1000; i++ { // all matches
		s1 := genInput(i + 1)
		s2 := genInput(i + 1)
		actual := MatchFold(s1, s2, 0, 0)
		assert.True(t, actual, fmt.Sprintf("failed for length: %v", 1+i))
	}
	for i := 0; i < 1000; i++ {
		s1 := genInput(i + 2)
		for j := 0; j < i; j++ {
			s2 := genInput(i + 2)
			s2[j] += 2
			actual := MatchFold(s1, s2, 0, 0)
			assert.False(t, actual, fmt.Sprintf("failed for not match: %v", 2+i))
		}
	}
}

var benchMatch = genInput(1000)

func BenchmarkMatchFoldFragment(b *testing.B) {
	for k := 0; k < b.N; k++ {
		actual := MatchFold(benchMatch, benchMatch, 0, 0)
		assert.True(b, actual)
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for k := 0; k < b.N; k++ {
		actual := bytes.EqualFold(benchMatch, benchMatch)
		assert.True(b, actual)
	}
}

func TestNewFragment(t *testing.T) {
	useCases := []struct {
		description string
		fragments   string
		options     []Option
		input       []byte
		matched     bool
	}{
		{
			description: "FragmentsFold exact match",
			fragments:   "xyz",
			input:       []byte("xyz"),
			matched:     true,
		},
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
		{
			description: "input same length as match fragment",
			fragments:   "abc",
			input:       []byte("abc"),
			matched:     true,
		},
	}

	for _, useCase := range useCases {
		matcher := NewFragment(useCase.fragments, useCase.options...)
		matched := matcher.Match(parsly.NewCursor("", useCase.input, 0))
		assert.Equal(t, useCase.matched, matched > 0, useCase.description)
	}

}
