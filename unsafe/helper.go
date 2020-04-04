package unsafe

import "unsafe"

//Converts bytes to string
func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

