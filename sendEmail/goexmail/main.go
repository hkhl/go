package main

import (
	"goexmail/controller"
)

func main() {
	//controller.SendEmail()
	//接口失效的时候放置默认数据

	today, tomorrow := controller.GetWeather("xian")
	poem := controller.GetOneSentencePoem()

	poemEx := "" + poem + "\r\n\n"
	WeathEx := today.Date + "  农历" + today.Lunar + "  " + today.Week + "\r\n" +
		"当前地区：" +today.Cityname + "      白天：" + today.WeatherDay + "           夜间：" + today.WeatherNight + "\r\n" +
		"最高温度：" + today.HighTemp + "℃       最低温度：" + today.LowTemp + "℃    风力等级：" + today.WindScale + "\r\n" +
		"明日：" + tomorrow.WeatherDay


	content := poemEx + WeathEx
	controller.SendEmail(content, "天气预报")


	//html邮件参考： https://www.cnblogs.com/hxlasky/p/11281957.html
}
