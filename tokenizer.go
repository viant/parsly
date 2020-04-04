package parsly

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

//Tokenizer represents a token scanner.
type Tokenizer struct {
	registry   *Registry
	input      []byte
	cursor     int
	location   *Location
	match      *Match
	invalid    *Token
	endOfFile  *Token
	whitespace *Token
}

//Cursor returns tokenizer cursor
func (t *Tokenizer) Cursor() int {
	return t.cursor
}

//NewError returns error for expected tokens
func (t *Tokenizer) NewError(expectedTokens ...int) error {
	var names = []string{}
	for _, expected := range expectedTokens {
		if token, ok := t.registry.Lookup(expected); ok {
			names = append(names, token.Name)
		}
	}
	return errors.Errorf("invalid JSON, expected: [%v] at  pos: %v", strings.Join(names, ","), (t.location.Offset + t.cursor))
}

//MatchAfterWhitespace matches token after whitespace
func (t *Tokenizer) MatchAfterWhitespace(whitespaceRequired bool, candidates ...int) *Match {
	matched := t.whitespace.Match(t.input, t.cursor)
	if whitespaceRequired && matched == 0 {
		return t.SetMatch(t.invalid, t.cursor, 0)
	}
	t.cursor += matched
	return t.MatchAny(candidates...)
}

//MatchAny matches the first of the candidates
func (t *Tokenizer) MatchAny(candidates ...int) *Match {
	var match *Match
	if t.cursor >= len(t.input) {
		return t.SetMatch(t.endOfFile, t.cursor, 0)
	}
	for _, candidate := range candidates {
		match = t.matchCandidate(candidate)
		if match.Token.Code == candidate {
			return match
		}
	}
	return match
}

func (t *Tokenizer) SetMatch(token *Token, cursor, matchSize int) *Match {
	t.match.Token = token
	t.match.Cursor = cursor
	t.match.Size = matchSize
	return t.match
}

//Match tries to match a candidate, it returns a match.
func (t *Tokenizer) Match(candidate int) *Match {
	cursor := t.cursor
	if !(cursor < len(t.input)) || candidate == t.endOfFile.Code {
		return t.SetMatch(t.endOfFile, cursor, 0)
	}
	return t.matchCandidate(candidate)
}

func (t *Tokenizer) matchCandidate(candidate int) *Match {
	cursor := t.cursor
	if token, ok := t.registry.Lookup(candidate); ok {
		matchedSize := token.Match(t.input, cursor)
		if matchedSize > 0 {
			result := t.SetMatch(token, cursor, matchedSize)
			t.cursor += matchedSize
			return result
		}
	} else {
		panic(fmt.Sprintf("failed to lookup token for %v", candidate))
	}
	return t.SetMatch(t.invalid, cursor, 0)
}


//NewTokenizer creates a new NewTokenizer,  token registry.
func NewTokenizer(location *Location, whitespace *Token, input []byte, registry *Registry) *Tokenizer {
	return &Tokenizer{
		registry:   registry,
		input:      input,
		invalid:    InvalidToken,
		endOfFile:  EOFToken,
		whitespace: whitespace,
		location:   location,
		match:      &Match{Location: location},
	}
}
