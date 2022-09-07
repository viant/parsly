package matcher

//IsLetter returns true if byte is ASCII letter
func IsLetter(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') {
		return false
	}
	return true
}
