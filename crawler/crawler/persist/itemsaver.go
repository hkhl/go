package persist

import (
	"crawler/engine"
	"errors"
	"gopkg.in/mgo.v2"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			// 接收到发送的 item
			item := <-out
			log.Printf("Item Saver: got item #%d: %v\n", itemCount, item)
			itemCount++

			// Save data in mongodb
			err := mongoSave(session, index, item)

			if err != nil {
				// if have err, ignore it
				log.Printf("Item Saver: error, saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

// 使用 MongoDB 保存数据
func mongoSave(session *mgo.Session, dbName string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	c := session.DB(dbName).C(item.Type)    // 选择要操作的数据库与集合
	err := c.Insert(item)        // 插入数据
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

