package parsly_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher"
	"testing"
)

func TestCursor_NewError(t *testing.T) {

	var useCases = []struct {
		description string
		input       []byte
		token       *parsly.Token
		expect      string
		hasMatch    bool
	}{
		{
			description: "error number",
			input:       []byte("aafdsfds"),
			token:       parsly.NewToken(1, "number", matcher.NewNumber()),
			expect:      "invalid token around *a*afdsfd,  expected: [number] at  pos: 0",
		},
	}

	for _, useCase := range useCases {
		cursor := parsly.NewCursor("", useCase.input, 0)
		err := cursor.NewError(useCase.token)
		if !assert.NotNil(t, err, useCase.description) {
			continue
		}
		assert.EqualValues(t, useCase.expect, err.Error(), useCase.description)
	}

}
func TestCursor_FindMatch(t *testing.T) {

	var useCases = []struct {
		description string
		input       []byte
		token       *parsly.Token
		expect      string
		hasMatch    bool
	}{
		{
			description: "find number",
			input:       []byte("abc123srx"),
			token:       parsly.NewToken(1, "number", matcher.NewNumber()),
			expect:      "123",
			hasMatch:    true,
		},
		{
			description: "find number",
			input:       []byte("zfgasdadasdas"),
			token:       parsly.NewToken(1, "number", matcher.NewNumber()),
			hasMatch:    false,
		},
	}

	for _, useCase := range useCases {
		cursor := parsly.NewCursor("", useCase.input, 0)
		match := cursor.FindMatch(useCase.token)
		if !useCase.hasMatch {
			assert.True(t, match.Code != useCase.token.Code, useCase.description)
			continue
		}
		assert.Equal(t, match.Code, useCase.token.Code, useCase.description)
		actual := match.Text(cursor)
		assert.EqualValues(t, useCase.expect, actual)
	}
}
