package parsly

//Token represents token matcher
type Token struct {
	Code int
	Name string
	Matcher
}

//NewToken creates a token
func NewToken(code int, name string, matcher Matcher) *Token {
	return &Token{
		Code:    code,
		Name:    name,
		Matcher: matcher,
	}
}
