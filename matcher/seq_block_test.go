package matcher

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/parsly"
	"testing"
)

func TestNewSeqBlock(t *testing.T) {

	useCases := []struct {
		description string
		begin       string
		end         string
		pos         int
		input       []byte
		matched     int
	}{

		{
			description: "match expr block",
			begin:       "#if",
			end:         "#end",
			input:       []byte("x#if ($Has.X) $X #endvc"),
			pos:         1,
			matched:     20,
		},
		{
			description: "match block",
			begin:       "BEGIN",
			end:         "END",
			input:       []byte("xBEGIN test ENDvc"),
			pos:         1,
			matched:     14,
		},
		{
			description: "sql comment",
			begin:       "/*",
			end:         "*/",
			input:       []byte(`/*{"Selector": {"Constraints": {"Projection": true, "Filterable": ["*"]}}}*/`),
			matched:     76,
		},
	}

	//for _, useCase := range useCases[len(useCases)-1:] {
	for _, useCase := range useCases {
		matcher := NewSeqBlock(useCase.begin, useCase.end)
		cursor := &parsly.Cursor{Input: useCase.input, Pos: useCase.pos}
		matched := matcher.Match(cursor)
		assert.Equal(t, useCase.matched, matched, useCase.description)
	}

}
