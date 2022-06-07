package controllers
//
//import (
//	"3TrashGuitarists/models"
//	"3TrashGuitarists/utils"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//)
//
//type SignInController struct {
//	beego.Controller
//}
//
//func (c *SignInController) Get() {
//	c.TplName = "Sign_In.html"
//
//}
//
//func (c *SignInController) Post() {
//	username := c.GetString("username")
//	password := c.GetString("password")
//
//	md5_pwd := utils.GetMd5(password)
//
//	o := orm.NewOrm()
//	exist := o.QueryTable(new(models.UserInformation3TG)).Filter("user_name",username).Filter("password",md5_pwd).Exist()
//
//	if exist == true{
//		c.SetSession("user_name",username)
//		c.Data["json"] = map[string]interface{}{"code": 1, "message": "Login Succeeded"}
//		c.Redirect(beego.URLFor("/"),302)
//	} else {
//		c.Data["json"] = map[string]interface{}{"code": 0, "message": "Login Failed"}
//	}
//	c.ServeJSON()
//
//}
