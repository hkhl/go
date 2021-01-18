package controller

import (
	"fmt"
	"goexmail/authorization"
	"goexmail/utils"
	"io/ioutil"
	"net/http"
)

func poem(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-User-Token", authorization.TokenOfPoem())

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	return string(s)
}

func GetOneSentencePoem() string{
	body := poem("https://v2.jinrishici.com/sentence")
	data := utils.String2Map(body)
	retStr := ""
	if (data["status"]) == "success" {
		sentence := data["data"].(map[string]interface{})["content"]   //考虑直接将json数据存储到结构体中 http://codingdict.com/questions/63399
		origin := data["data"].(map[string]interface{})["origin"]
		author := origin.(map[string]interface{})["author"]
		retStr = fmt.Sprintf("%s ——%s", sentence, author)
	}
	return retStr
}