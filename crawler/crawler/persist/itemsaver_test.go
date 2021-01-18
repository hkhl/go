// persist/itemsaver_test.gp
package persist

import (
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)
func TestMongoSave(t *testing.T) {
	// mongodb connect
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1946858930",
		Type: "zhenai",
		Id:   "1946858930",
		Payload: model.Profile{
			Name:     "為你垨候",
			Gender:   "女士",
			Age:      40,
			Height:   163,
			Weight:   54,
			Income:   "5-8千",
			Marriage: "未婚",
			Address:  "佛山顺德区",
		},
	}
	// 保存数据
	err = mongo_save(session, "crawler", expected)
	if err != nil {
		panic(err)
	}

	c := session.DB("crawler").C("zhenai")

	var result engine.Item
	// 查询数据
	err = c.Find(bson.M{"id": "1946858930"}).One(&result)
	// result 为 Json 类型
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s, %s, %v\n", result.Url, result.Id, result.Payload)
}