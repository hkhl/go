package model

import "encoding/json"

//用户个人信息
type Profile struct {
	//Url      string
	//Id       string
	Name     string
	Gender   string
	Age      int
	Height   int
	Weight   int
	Income   string
	Marriage string
	Address  string
}

func FormJsonObj(obj interface{}) (Profile, error) {
	var profile Profile
	string, err := json.Marshal(obj)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(string, &profile)
	return profile, err
}
