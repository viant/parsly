package parsly

import (
	"bytes"
	"strconv"
)

// TokenMatch represents a token match
type TokenMatch struct {
	Offset int
	Size   int
	*Token
}

// Matched return matched fragment
func (m *TokenMatch) Bytes(cursor *Cursor) []byte {
	return cursor.Input[m.Offset : m.Offset+m.Size]
}

// MatchedUnquoted return matched unquoted fragment
func (m *TokenMatch) UnquotedText(cursor *Cursor) string {
	data := m.Bytes(cursor)
	return AsString(data[1 : len(data)-1])
}

// Contains return true if lastMatch data contains rune
func (m *TokenMatch) ContainsRune(cursor *Cursor, r rune) bool {
	return bytes.ContainsRune(cursor.Input[m.Offset:m.Offset+m.Size], r)
}

// Matched return matched fragment
func (m *TokenMatch) Text(cursor *Cursor) string {
	data := m.Bytes(cursor)
	return AsString(data)
}

// Byte return matched byte
func (m *TokenMatch) Byte(cursor *Cursor) byte {
	data := m.Bytes(cursor)
	return data[0]
}

// Float return matched float value
func (m *TokenMatch) Float(cursor *Cursor) (float64, error) {
	data := m.Bytes(cursor)
	value := AsZeroAllocString(data)
	return strconv.ParseFloat(value, 64)
}

// Int return matched int value
func (m *TokenMatch) Int(cursor *Cursor) (int64, error) {
	data := m.Bytes(cursor)
	value := AsZeroAllocString(data)
	return strconv.ParseInt(value, 10, 64)
}

// Uint return matched uint value
func (m *TokenMatch) Uint(cursor *Cursor) (uint64, error) {
	data := m.Bytes(cursor)
	value := AsZeroAllocString(data)
	return strconv.ParseUint(value, 10, 64)
}

// Bool return matched bool value
func (m *TokenMatch) Bool(cursor *Cursor) (bool, error) {
	data := m.Bytes(cursor)
	value := AsZeroAllocString(data)
	return strconv.ParseBool(value)
}

// SetToken sets token
func (m *TokenMatch) SetToken(token *Token, cursor int, size int) {
	m.Token = token
	m.Offset = cursor
	m.Size = size
}
