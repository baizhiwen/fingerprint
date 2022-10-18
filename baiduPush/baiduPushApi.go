package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/**
//"https://www.lanjinger.com/d/192647",
		//"https://www.lanjinger.com/d/192684",
		//"https://www.lanjinger.com/d/192436",
		//"https://www.lanjinger.com/d/192762",
		//"https://www.lanjinger.com/d/192680",
		//"https://www.lanjinger.com/d/192619",
		//"https://www.lanjinger.com/d/192528",
https://www.lanjinger.com/d/192526
https://www.lanjinger.com/d/192476
https://www.lanjinger.com/d/192808
https://www.lanjinger.com/d/192809
https://www.lanjinger.com/d/192799
https://www.lanjinger.com/d/192798
https://www.lanjinger.com/d/192791
https://www.lanjinger.com/d/192774
https://www.lanjinger.com/d/192783
https://www.lanjinger.com/d/192782
https://www.lanjinger.com/d/192694
https://www.lanjinger.com/d/192659
 */
func main() {
	ans := PushApi("http://data.zz.baidu.com/urls?site=https://www.lanjinger.com&token=",
		"https://www.lanjinger.com/d/192659",
		)
	fmt.Println(ans)
}

func PushApi(url string, data string) string {
	/*        百度推送API
				发送POST请求
				url：         请求地址：http://data.zz.baidu.com/urls?site=https://www.helloworldtools.com&token=xxxx
				data：        POST请求提交的数据(推送的连接)
				contentType： 请求体格式: text/plain
				成功返回示例
				{
					"remain":99998,
					"success":2,
					"not_same_site":[],
					"not_valid":[]
				}*/
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, "text/plain", strings.NewReader(data))
	resp.Header.Add("User-Agent", "curl/7.12.1")
	resp.Header.Add("Host", "data.zz.baidu.com")
	resp.Header.Add("Content-Type", "text/plain")
	//resp.Header.Add("Content-Length", "83")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}