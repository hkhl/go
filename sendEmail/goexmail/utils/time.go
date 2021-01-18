package utils

import (
	"github.com/nosixtools/solarlunar"
	"time"
)

//day=-1 昨天
//day=0  今天
//day=1  明天
func Weekday(day int) string {
	t := time.Now()
	retTime := t.AddDate(0, 0, day)
	return retTime.Weekday().String()
}

func SolarToChineseLuanr(solarDate string) string {
	//fmt.Println(solarlunar.SolarToChineseLuanr(solarDate))
	//fmt.Println(solarlunar.SolarToSimpleLuanr(solarDate))
	return solarlunar.SolarToChineseLuanr(solarDate)
}
