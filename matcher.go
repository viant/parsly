package parsly

//Matcher represents a matcher, that matches input from offset position, it returns number of bytes matched.
type Matcher interface {
	//TokenMatch matches input starting from offset, it return number of characters matched
	Match(cursor *Cursor) (matched int)
}
