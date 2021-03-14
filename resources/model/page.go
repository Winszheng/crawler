package model

import "github.com/Winszheng/crowler/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
