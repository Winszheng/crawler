package controller

import (
	"fmt"
	"github.com/Winszheng/crowler/resources/model"
	"github.com/Winszheng/crowler/resources/view"
	"github.com/olivere/elastic/v7"
	"net/http"
	"strconv"
	"strings"
)

// SearchResultHandler从elastic获取数据给view
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// 给文件名才能
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	// 连上elasticsearch，建立好view
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from")) // 分页
	if err != nil {                                  // 出错或用户乱输入
		from = 0
	}

	var page model.SearchResult
	page = getSearchResult(q, from)
	err = h.view.Render(w, page)
	if err != nil {
		http.Error()
	}
}
