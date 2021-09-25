package main

import (
	"github.com/Winszheng/crawler/af/resources/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("resources/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"resources/list.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
