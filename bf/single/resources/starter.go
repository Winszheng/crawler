package main

import (
	"github.com/Winszheng/crawler/single/resources/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("single/resources/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"single/resources/list.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
