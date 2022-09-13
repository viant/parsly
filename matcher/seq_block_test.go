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
			description: "block with asterisks, and odd number of quotation marks (1)",
			begin:       "*",
			end:         "*",
			input:       []byte(`*"*`),
			pos:         0,
			matched:     3,
		},
		{
			description: "block with asterisks, and even number of quotation marks (2)",
			begin:       "*",
			end:         "*",
			input:       []byte(`*""*`),
			pos:         0,
			matched:     4,
		},
		{
			description: "block with asterisks, and odd number of quotation marks (3)",
			begin:       "*",
			end:         "*",
			input:       []byte(`*"""*`),
			pos:         0,
			matched:     5,
		},
		{
			description: "block with asterisks, and odd number of quotation marks (3) and additional spaces",
			begin:       "*",
			end:         "*",
			input:       []byte(`* " " " *`),
			pos:         0,
			matched:     9,
		},
		{
			description: "block with asterisks, and even number of quotation marks (4)",
			begin:       "*",
			end:         "*",
			input:       []byte(`*""""*`),
			pos:         0,
			matched:     6,
		},
		{
			description: "block with asterisks, and even number of quotation marks (4) and additional spaces",
			begin:       "*",
			end:         "*",
			input:       []byte(`* " " " " *`),
			pos:         0,
			matched:     11,
		},
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
