package engine

// request和item是engine用的，所以放在engine这里

type Request struct {
	Url       string
	ParseFunc func([]byte, string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Iterms   []Item
}

type Item struct {
	Url      string
	Id       string
	Playload interface{}
}

// NilParser is a palceholder
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
