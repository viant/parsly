package parsly

const (
	//EOF end of file token
	EOF = -1
	//Invalid invalid token
	Invalid = -2
)

//EOFToken
var EOFToken = NewToken(EOF, "EOF", nil)

//InvalidToken represents invalid token
var InvalidToken = NewToken(Invalid, "Invalid", nil)

