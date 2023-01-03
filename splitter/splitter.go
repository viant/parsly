package splitter

import (
	"github.com/viant/parsly"
	"github.com/viant/parsly/matcher"
	"strconv"
)

type Splitter struct {
	tokens []*parsly.Token
}

func (s *Splitter) Split(text string) []string {
	var result = make([]string, 0)
	cursor := parsly.NewCursor("", []byte(text), 0)
	limit := cursor.Pos
outer:
	for {
		match := cursor.MatchAny(s.tokens...)
		switch match.Code {
		case parsly.Invalid:
			cursor.Pos++
		case parsly.EOF:
			item := cursor.Input[limit:]
			result = append(result, string(item))
			break outer
		default:
			offset := cursor.Pos - match.Size
			item := cursor.Input[limit:offset]
			result = append(result, string(item))
			limit = cursor.Pos
		}
	}
	return result
}

//New creates a splitter with supplied delimiters, it takes Case option to control case sensivity
func New(delimiters []string, options ...matcher.Option) *Splitter {
	var tokens []*parsly.Token
	for i := range delimiters {
		delimiterMatcher := matcher.NewFragment(delimiters[i], options...)
		tokens = append(tokens, &parsly.Token{Code: i, Name: "DEL" + strconv.Itoa(i), Matcher: delimiterMatcher})
	}
	return &Splitter{tokens: tokens}
}
