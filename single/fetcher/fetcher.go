package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch 给定url获取文档
func Fetch(url string) ([]byte, error){

	//const url = `http://album.zhenai.com/u/1883184587`
	//const url1 = `http://m.zhenai.com/u/1275335590`
	// 前者有反爬机制
	// url = strings.Replace(url, "album", "m", -1)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")

	// 这个cookie只有一分钟
	// 这也太鬼了
	//cookie := "sid=23d5de3c-b488-4cb5-a755-227ebf836fc3; ec=a7JXiQOh-1613308803385-f7ea684eea747-135889905; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1613308811; FSSBBIl1UgzbN7NO=52ElueikvfOhj5hTFXBuUm5BazpXF1KPcvrhVgHcPx1ccb4YoDkmb1yUh36tJwLWOZuV49wlAgDa9PM8G.NmTvq; spliterabparams=1613569148894:6149511542232997122; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1613571045; FSSBBIl1UgzbN7NP=53c0OyCqzE5Aqqqm67mpl3GxNhLCiG.uYz4JGTjb.69l.5hU9Xy0oLMHvPGBo01s5BCNU8gUjqz2wsS8wQBgzGfTM5KSlJEuApBPxx8CpT5iUzqCV0MFzrt3ueuIEUWa17lgi91_eK6n6RNmSBKRCddcSKp51uNgFPWNqP9tTHAnslRNMzb5Uv4RPnNhX.j8B.LmBhODBjrKRw5ZwLwB7Y4jpI8NX2kqnVyYlxGnMOsfex6h2T0P.mkHzusRQmWEqa; _efmdata=b3jiMWME3aliz6py8qy26e4+N4pwmXOY4czMFZKJo04XAtd4/TzdC+Hb/uCBSoUIRGjgs1yI/HaRFrmrYodxVCKb6tAfd+nWk6QgvmySU54=; _exid=fIIHmsFyVuFUioGmgc9QOcT2LOkqFRtTDi+XTuVq/oepy63mEhaeomJ5W93nJHoGA4TnjOdr16Ff96oQgab2ew=="
	//req.Header.Add("cookie", cookie)

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
