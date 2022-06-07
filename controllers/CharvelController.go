package controllers

type CharvelController struct {
	//beego.Controller
	BaseController
}

func (c *CharvelController) Get() {
	c.TplName = "Brand_Charvel.html"

}
