package ext

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// Post 执行post请求
func Post(urlObj *url.URL, params map[string]string) string {
	urlParams := url.Values{}
	for k, v := range params {
		urlParams.Set(k, v)
	}
	urlObj.RawQuery = urlParams.Encode()
	urlPath := urlObj.String()
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
