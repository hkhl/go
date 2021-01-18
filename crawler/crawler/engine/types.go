package engine

type Request struct {
	Url	 		string
	ParseFunc 	func([]byte) ParseResult //解析函数
}

type ParseResult struct {
	Requests 	[]Request
	Items 		[]Item
	//Items 		[]interface{}
}

// 解析出的用户数据格式
type Item struct {
	Url     string      // 个人信息Url地址
	Type    string      // table
	Id      string      // Id
	Payload interface{} // 详细信息
}