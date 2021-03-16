package view

import (
	"github.com/Winszheng/crowler/resources/model"
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
}

// Render用模板和数据渲染出前端
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
