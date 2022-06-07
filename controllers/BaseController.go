package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	user_name interface{}
}

func (c *BaseController) Prepare() {
	user_name := c.GetSession("user_name")
	fmt.Println("user_name---->", user_name)
	if user_name != nil {
		c.IsLogin = true
		c.user_name = user_name
	} else {
		c.IsLogin = false
	}
	c.Data["IsLogin"] = c.IsLogin
	c.Data["user_name"] = c.user_name
}