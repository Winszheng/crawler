package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch 给定url获取文档
// share
func Fetch(url string) ([]byte, error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// 返回error的方法
		// 1.用fmt.Errorf生成一个error
		// 2.用errors.New生成error
		// 3.自己写一个struct, 实现Error() string即可
		return nil, fmt.Errorf("error: status code %v", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// 直接用http.get遇到了403错误
// 原因：在短时间内直接get获取大量数据，会被服务器认为是攻击，因此请求会被拒绝。
// 解决方法：
// 1.包装请求，伪装成浏览器请求模式，浏览器一次获取大量数据并不罕见，如果是程序一次获取大量数据就有些可疑。但有时服务器是根据同一ip请求频率来判断的，因此即使伪装成不同的浏览器也不会奏效。（ip被封）
// 2.降低请求频率，比如设置0.1s的暂停时间。
// 3.使用代理ip，即通过自动更换不同ip来欺骗服务器，让它认为是来自不同电脑的访问请求，从而不会被拒绝。代理ip可以在某些网站寻找。
// 4.终极操作：把更换不同header和更换ip结合。组合数越多越好。

// ddos攻击：短时间内发起大量请求，耗尽服务器资源，以至于无法响应正常访问，造成网站实质性下线。