package controllers

import (
	"3TrashGuitarists/models"
	"3TrashGuitarists/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SignInController struct {
	beego.Controller
}

func (c *SignInController) Get() {
	c.TplName = "Sign_In.html"

}

func (c *SignInController) Post() {
	ReturnData := make(map[string]interface{})

	username :=c.Input().Get("username")
	password :=c.Input().Get("password")

	md5_pwd := utils.GetMd5(password)
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(md5_pwd)

	o := orm.NewOrm()
	exist := o.QueryTable(new(models.UserInformation3TG)).Filter("user_name",username).Filter("password",md5_pwd).Exist()

	if exist == true{
		ReturnData["status"] =1
		ReturnData["message"] ="Login Succeeded"
		c.SetSession("user_name",username)
		//c.Redirect(beego.URLFor("/"),302)
	} else {
		ReturnData["status"] =0
		ReturnData["message"] ="Login failed. Please check your account name and password."

	}
	c.Data["json"] = ReturnData
	c.ServeJSON()

	fmt.Println(ReturnData)
}
