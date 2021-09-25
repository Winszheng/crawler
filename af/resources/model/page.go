package model

// 信息模型
// 在html页面的作用相当于占位符
type SearchResult struct {
	Hits  int
	Start int
	Items []interface{}
}
