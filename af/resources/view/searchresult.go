package view

import (
	"github.com/Winszheng/crawler/af/resources/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

// CreateSearchResultView返回SearchResultView(模板结构体)
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{template: template.Must(
		template.ParseFiles(filename))}
	// 把对应文件解析成*Template, 把template.ParseFiles和template.Must一起使用，是为了出错直接panic，省事
}

// Render把data渲染到模板w
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
