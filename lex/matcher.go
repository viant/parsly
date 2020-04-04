package lex


//Matcher represents a matcher, that matches input from offset position, it returns number of bytes matched.
type Matcher interface {
	//Match matches input starting from offset, it return number of characters matched
	Match(input []byte, offset int) (matched int)
}

//RuneMatcher represents a rune matcher
type RuneMatcher interface {
	MatchRune(runeValue rune) bool
}