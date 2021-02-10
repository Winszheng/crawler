package engine

// 这里定义一些会用到的类型

// 这个项目还是在带你用golang啊
// ccmouse牛逼！

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult // 成员变量是个函数
}

// 果然，切片比数组好用啊
// 空切片yyds
type ParseResult struct {
	// 空切片
	Requests []Request
	// 空切片
	Items    []interface{} //interface{}: 空接口表任何类型==>暂时不规定Items类型==>实际上Items表城市名或用户名
}

// NilParser相当于占位符
// 当Parser还没写好时，为了编译正常占位，返回空集合
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
