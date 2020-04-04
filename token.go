package parsly

import "github.com/viant/parsly/lex"

//Token represents token matcher
type Token struct {
	Code int
	Name string
	lex.Matcher
}

//NewToken creates a token
func NewToken(code int, name string, matcher lex.Matcher) *Token {
	return &Token{
		Code:    code,
		Name:    name,
		Matcher: matcher,
	}
}
