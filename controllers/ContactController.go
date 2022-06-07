package controllers

type ContactController struct {
	//beego.Controller
	BaseController
}

func (c *ContactController) Get() {
	c.TplName = "Contact.html"

}

