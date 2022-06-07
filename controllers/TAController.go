package controllers

type TAController struct {
	//beego.Controller
	BaseController
}

func (c *TAController) Get() {
	c.TplName = "Brand_TA.html"

}
