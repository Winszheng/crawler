package engine

// 这里定义一些会用到的类型

// 这个项目还是在带你用golang啊
// ccmouse牛逼！

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult // url对应的解析函数
}

// 果然，切片比数组好用啊
// 空切片yyds
type ParseResult struct {
	Requests []Request
	Items    []interface{} //interface{}: 空接口表任何类型==>暂时不规定Items类型==>实际上Items表城市名或用户名
}

// NilParser相当于占位符
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
