package fetcher

import (
	"fmt"
	"github.com/Winszheng/crawler/af/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// timer 用于避免发请求速度太快，触发反爬机制
// 0.2s请求一次
var timer = time.Tick(time.Second / config.Qps)

// Fetcher fetches contents from specific url
// return contents and error
func Fetcher(url string) ([]byte, error) {
	<-timer
	log.Println("Fetching url:", url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // 状态码200，可以获取信息
		// 要返回error，生成error的3个办法：
		// 1.errors.New("error msg")
		// 2.fmt.Errorf
		// 3.自己实现error接口
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
