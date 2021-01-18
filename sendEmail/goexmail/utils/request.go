package utils

import (
	"io/ioutil"
	"net/http"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	return string(s)
}

////  http 简单封装参考 ： https://www.cnblogs.com/zhaof/p/11346412.html
