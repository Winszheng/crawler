package fetcher

import (
	"testing"
)

const url = `http://album.zhenai.com/u/1883184587`

func TestFetch(t *testing.T) {
	Fetch(url)

}