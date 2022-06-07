package controllers

type SuhrController struct {
	//beego.Controller
	BaseController
}

func (c *SuhrController) Get() {
	c.TplName = "Brand_Suhr.html"

}
