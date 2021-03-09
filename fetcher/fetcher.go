package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 为了规避反爬机制
var rateLimiter = time.Tick(100 * time.Millisecond) // ==> 0.1s

// Sleep是使用睡眠完成定时，结束后继续往下执行循环来实现定时任务。
// Tick函数是使用channel阻塞当前协程，完成定时任务的执行

func Fetch(url string) ([]byte, error) {
	<-rateLimiter // 阻塞避免向server请求数据的频率过快
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
