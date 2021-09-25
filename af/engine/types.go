package engine

// 这不就是变相构造除了链表吗？
type Request struct {
	Url string
	Parser func([]byte, string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}	// 因为这个Item实际上只有用户个人信息那块，才能用上，所以我觉得也没必要emmm
}

// 空parser，相当于占位符
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
