package model

// 信息模型
type SearchResult struct {
	Hits  int
	Start int
	Items []interface{}
}
