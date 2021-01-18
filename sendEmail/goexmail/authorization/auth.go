package authorization

import (
	"goexmail/utils"
)

func TokenOfPoem() string {
	body := utils.Get("https://v2.jinrishici.com/token")
	bodyMap := utils.String2Map(body)

	if bodyMap["status"] == "success" {
		return bodyMap["data"].(string)
	} else {
		panic("token获取失败")
	}

	// 三种类型转换 https://studygolang.com/articles/21591?fr=sidebar
	// 1、普通变量  int(var)
	// 2、指针变量  (*int)(unsafe.Pointer(var))
	// 3、类型断言  var.(int)
}
