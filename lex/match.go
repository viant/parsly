package lex



//MatchRune matcher rune with matchers, it reutnr the first matched matcher or nil
func MatchRune(runeValue rune,  matchers ... RuneMatcher) RuneMatcher {
	for _, matcher := range matchers {
		if matcher.MatchRune(runeValue) {
			return matcher
		}
	}
	return nil
}
