package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：921210335"
	c.Data["qq"] = "QQ：921210335"
	c.Data["tel"] = "Tel：15619292231"
	c.TplName = "aboutme.html"
}

