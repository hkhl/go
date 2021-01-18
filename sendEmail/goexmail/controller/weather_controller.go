package controller

import (
	"fmt"
	"goexmail/model"
	"goexmail/utils"
)


//根据城市明，返回今天和明天的天气
func GetWeather(city string) (today, tomorrow model.WeatherModle) {
	var weatherModel [2]model.WeatherModle
	city = "xian"
	url := fmt.Sprintf("https://api.seniverse.com/v3/weather/daily.json?key=S5sq3dwhu6iF7OHks&location=%s&language=zh-Hans&unit=c&start=0&days=2", city)
	rsp := utils.Get(url)
	rspMap := utils.String2Map(rsp)

	results := rspMap["results"].([]interface{})[0]
	location := results.(map[string]interface{})["location"]
	daily := results.(map[string]interface{})["daily"].([]interface{})

	for i := 0; i <= 1; i++ {
		tmp :=  daily[i].(map[string]interface{})
		weatherModel[i].Date = tmp["date"].(string) //tmp["date"].(string)
		weatherModel[i].Week = utils.Weekday(i)
		weatherModel[i].Lunar = utils.SolarToChineseLuanr(weatherModel[i].Date)
		weatherModel[i].Cityname = location.(map[string]interface{})["name"].(string)
		weatherModel[i].WeatherDay = tmp["text_day"].(string)
		weatherModel[i].WeatherNight = tmp["text_night"].(string)
		weatherModel[i].HighTemp = tmp["high"].(string)
		weatherModel[i].LowTemp = tmp["low"].(string)
		weatherModel[i].WindScale = tmp["wind_scale"].(string)
		weatherModel[i].WindSpeed = tmp["wind_speed"].(string)

	}
	//fmt.Println(weatherModel[0], weatherModel[1])
	return weatherModel[0], weatherModel[1]
}
