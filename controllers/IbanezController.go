package controllers

type IbanezController struct {
	//beego.Controller
	BaseController
}

func (c *IbanezController) Get() {
	c.TplName = "Brand_Ibanez.html"

}
