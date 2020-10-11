package parsly

import (
	"github.com/pkg/errors"
	"strings"
)

//Cursor represents a location
type Cursor struct {
	Path      string
	offset    int
	Input     []byte
	InputSize int
	Pos       int
	lastMatch TokenMatch
}

//NewError returns error for expected tokens
func (c *Cursor) NewError(expectedTokens ...*Token) error {
	var names = []string{}
	for _, token := range expectedTokens {
		names = append(names, token.Name)
	}
	return errors.Errorf("invalid token, expected: [%v] at  pos: %v", strings.Join(names, ","), c.offset+c.Pos)
}

//TokenMatch returns updated lastMatch
func (c *Cursor) TokenMatch(token *Token, matchSize int) *TokenMatch {
	c.lastMatch.Token = token
	c.lastMatch.Offset = c.Pos
	c.lastMatch.Size = matchSize
	return &c.lastMatch
}

//HasMore returns true if it has more
func (c *Cursor) HasMore() bool {
	return c.Pos < c.InputSize
}

//MatchAfterOptional matcher first candidate after optional token lastMatch
func (c *Cursor) MatchAfterOptional(optional *Token, candidates ...*Token) *TokenMatch {
	if !c.HasMore() {
		return c.TokenMatch(EOFToken, 0)
	}
	matched := optional.Match(c)
	c.Pos += matched
	return c.MatchAny(candidates...)
}

//MatchAny matches the first of the candidates
func (c *Cursor) MatchAny(candidates ...*Token) *TokenMatch {
	if !c.HasMore() {
		return c.TokenMatch(EOFToken, 0)
	}
	size := candidates
	i := 0
loop:
	token := candidates[i]
	matchedSize := token.Match(c)
	if matchedSize > 0 {
		result := c.TokenMatch(token, matchedSize)
		c.Pos += matchedSize
		return result
	}
	i++
	if i < len(size) {
		goto loop
	}
	return c.TokenMatch(InvalidToken, 0)
}

//MatchOne tries to lastMatch a candidate, it returns a lastMatch.
func (c *Cursor) MatchOne(token *Token) *TokenMatch {
	if !c.HasMore() {
		return c.TokenMatch(EOFToken, 0)
	}
	matchedSize := token.Match(c)
	if matchedSize > 0 {
		result := c.TokenMatch(token, matchedSize)
		c.Pos += matchedSize
		return result
	}
	return c.TokenMatch(InvalidToken, 0)
}


//FindMatch tries to find a token match in the cursor
func (c *Cursor) FindMatch(token *Token) *TokenMatch {
	pos := c.Pos
	for ; ; {
		if !c.HasMore() {
			break
		}
		matchedSize := token.Match(c)
		if matchedSize > 0 {
			result := c.TokenMatch(token, matchedSize)
			c.Pos += matchedSize
			return result
		}
		c.Pos++
	}
	//
	c.Pos = pos
	return c.TokenMatch(InvalidToken, 0)
}

//NewCursor creates a location
func NewCursor(path string, input []byte, offset int) *Cursor {
	return &Cursor{Path: path, Input: input, offset: offset, InputSize: len(input)}
}
