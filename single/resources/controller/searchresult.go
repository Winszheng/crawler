package controller

import (
	"context"
	"github.com/Winszheng/crawler/single/engine"
	"github.com/Winszheng/crawler/single/resources/model"
	"github.com/Winszheng/crawler/single/resources/view"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// SearchResultHandler从elastic从client获取数据给view
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// CreateSearchResultHandler配置相应的template和连接elasticsearch
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
	from, err := strconv.Atoi(req.FormValue("from")) // 分页，从第from条record开始
	if err != nil {                                  // 出错或用户乱输入
		from = 0
	}

	var page model.SearchResult
	page, err = h.getSearchResult(q, from) // get data
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = int(resp.TotalHits()) // 命中了几条记录
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	return result, err
}
