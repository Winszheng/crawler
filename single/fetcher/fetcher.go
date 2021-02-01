package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetch 给定url获取文档
func Fetch(url string) ([]byte, error){
	resp, err := http.Get(url) // seed
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// 返回error的方法
		// 1.用fmt.Errorf生成一个error
		// 2.用errors.New生成error
		// 3.自己写一个struct, 实现Error() string即可
		return nil, fmt.Errorf("error: status code ", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
