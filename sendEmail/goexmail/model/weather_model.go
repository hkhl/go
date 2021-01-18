package model

type WeatherModle struct {
	Date         string //时间
	Week         string //星期
	Lunar        string //农历时间
	Cityname     string //城市名
	WeatherDay   string //白天天气
	WeatherNight string //夜晚天气
	Temp         string //当前温度
	HighTemp     string //最高温度
	LowTemp      string //最低温度
	WindSpeed    string //风速km/h
	WindScale    string //风力等级
	Tips         string //小提示
}
