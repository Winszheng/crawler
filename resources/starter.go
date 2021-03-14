package main

import (
	"github.com/Winszheng/crowler/resources/controller"
	"net/http"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler(
		"resources/list.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}