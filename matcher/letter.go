package matcher

// IsLetter returns true if byte is ASCII letter
func IsLetter(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') {
		return false
	}
	return true
}

// IsDigit returns true if byte is ASCII digit
func IsDigit(b byte) bool {
	if b < '0' || b > '9' {
		return false
	}
	return true
}
