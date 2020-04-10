package matcher

//IsByte returns true if rune is byte
func isByte(value rune) bool {
	return rune(byte(value)) == value
}
