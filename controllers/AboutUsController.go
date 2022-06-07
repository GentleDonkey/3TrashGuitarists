package controllers

type AboutUsController struct {
	//beego.Controller
	BaseController
}

func (c *AboutUsController) Get() {
	c.TplName = "About_Us.html"

}