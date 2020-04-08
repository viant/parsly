package parsly


//AsString converts bytes to string
var AsString = func(data []byte) string {
	return string(data)
}

//AsZeroAllocString converts bytes to string,
//In this package I do not any unsafe dependency
//But if you need zero allocation override AsZeroAllocString =  usage.ByteSlice2String
var AsZeroAllocString = func(data []byte) string {
	return string(data)
}

