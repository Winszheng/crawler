package main

import (
	"fmt"
	"strings"
)

var url = `http://album.zhenai.com/u/1883184587`
const url1 = `http://m.zhenai.com/u/1275335590`

func main() {
	url = strings.Replace(url, "album", "m", -1)
	fmt.Println(url)
	//fmt.Println("hello")
	//_, err := fetcher.Fetch(url)
	//fmt.Println(err)
	//_, err = fetcher.Fetch(url1)
	//fmt.Println(err)
}
