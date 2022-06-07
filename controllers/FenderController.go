package controllers

type FenderController struct {
	//beego.Controller
	BaseController
}

func (c *FenderController) Get() {
	c.TplName = "Brand_Fender.html"

}