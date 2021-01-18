package utils

import "encoding/json"

// string转map
func String2Map(str string) map[string]interface{} {
	strMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &strMap)
	if err != nil {
		panic(err)
	}
	return strMap
}
