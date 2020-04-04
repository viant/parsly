package parsly


//Match represents
type Match struct {
	Location *Location
	Cursor   int
	Size     int
	*Token
}



//AsString converts bytes to string
var AsString = func(data []byte) string {
	return string(data)
}


//Matched return matched fragment
func (m *Match) Bytes(tokenizer *Tokenizer) []byte {
	return tokenizer.input[m.Cursor : m.Cursor+m.Size]
}





//MatchedUnquoted return matched unquoted fragment
func (m *Match) UnquotedText(tokenizer *Tokenizer) string {
	data := m.Bytes(tokenizer)
	return AsString(data[1:len(data)-1])
}


//Matched return matched fragment
func (m *Match) Text(tokenizer *Tokenizer) string {
	data := m.Bytes(tokenizer)
	return AsString(data)
}


func (m *Match) SetToken(token *Token, cursor int, size int) {
	m.Token = token
	m.Cursor = cursor
	m.Size = size
}

func NewMatch(location *Location) *Match {
	return &Match{Location:location}
}