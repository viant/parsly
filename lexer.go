package parsly

// Lexer is a struct that holds the tokens
type Lexer struct {
	tokens map[string]*Token
}

func NewLexer() *Lexer {
	return &Lexer{tokens: make(map[string]*Token)}
}

func (l *Lexer) Register(token *Token) {
	l.tokens[token.Name] = token
}

func (l *Lexer) Token(name string) *Token {
	return l.tokens[name]
}

func (l *Lexer) Tokens() map[string]*Token {
	return l.tokens
}
