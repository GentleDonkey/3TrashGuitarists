package controllers

type ESPController struct {
	//beego.Controller
	BaseController
}

func (c *ESPController) Get() {
	c.TplName = "Brand_ESP.html"

}

