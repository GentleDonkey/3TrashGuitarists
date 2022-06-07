package controllers

import "fmt"

type HomeController struct {
	//beego.Controller
	BaseController
}

func (c *HomeController) Get() {
	fmt.Println("IsLogin:",c.IsLogin,c.user_name)
	c.TplName = "index.html"

}
