package engine

type ParserFunc func([]byte, string) ParseResult

// request和item是engine用的，所以放在engine这里
// 为什么要单独包装出Parser这个接口呢？
// 啊好困惑
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser // 这个接口不匿名的意义在哪里？
	// 改动是把函数转变成接口
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
type NilParser struct{}

func (n NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// 工厂模式
// 因为FuncParser使用指针实现Parser接口，所以要返回地址来给Request的接口变量赋值
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
