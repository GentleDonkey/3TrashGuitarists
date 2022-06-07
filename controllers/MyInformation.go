package controllers

import (
	"3TrashGuitarists/models"
	"3TrashGuitarists/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MyInformationController struct {
	//beego.Controller
	BaseController
}

func (c *MyInformationController) Get() {
	user_name := c.GetSession("user_name")
	fmt.Println("user_name---->", user_name)
	if user_name != nil {
		c.IsLogin = true
		c.user_name = user_name
	} else {
		c.IsLogin = false
	}


	if c.GetSession("user_name") !=nil{
		c.IsLogin = true
		c.user_name = c.GetSession("user_name")
		o := orm.NewOrm()
		user := models.UserInformation3TG{}
		err := o.QueryTable(new(models.UserInformation3TG)).Filter("UserName",c.user_name).One(&user)
		fmt.Println(err)
		if err == nil{
			fmt.Println(user)
			c.Data["UserName"] = user.UserName
			c.Data["Email"] = user.Email
			c.Data["Address1"] = user.Address1
			c.Data["Address2"] = user.Address2
			c.Data["Address3"] = user.Address3
		}else {
			fmt.Println("data loading error")
		}

	}else{
		c.Redirect(beego.URLFor("/sign_in"),302)
	}
	c.TplName = "My_Information.html"

}

func (c *MyInformationController) ChangePassword() {
	password :=c.Input().Get("password")
	repassword :=c.Input().Get("repassword")
	ReturnData := make(map[string]interface{})

	if password != repassword{
		ReturnData["status"] = 0
		ReturnData["message"] = "2 passwords should be the same."
	}else if utils.VerifyPasswordFormat(password)==false{
		ReturnData["status"] = 0
		ReturnData["message"] = "Password should be in length of 6~16 with at least 1 character."
	}else{
		user := models.UserInformation3TG{}
		o := orm.NewOrm()
		o.QueryTable(new(models.UserInformation3TG)).Filter("UserName",c.GetSession("user_name")).One(&user)
		err := o.Read(&user)
		if err==nil{
			user.Password = utils.GetMd5(password)
			_, err1 := o.Update(&user)
			if err1==nil{
				ReturnData["status"] = 1
				ReturnData["message"] = "Password changed successfully."
				fmt.Println(user)
			}
		}
	}
	c.Data["json"] = ReturnData
	c.ServeJSON()
}

func (c *MyInformationController) ChangeAddress1() {
	address1 :=c.Input().Get("address1")
	ReturnData := make(map[string]interface{})

	if address1==""{
		ReturnData["status"] = 0
		ReturnData["message"] = "Please do not leave Address1 blank."
	} else{
		user := models.UserInformation3TG{}
		o := orm.NewOrm()
		o.QueryTable(new(models.UserInformation3TG)).Filter("UserName",c.GetSession("user_name")).One(&user)
		err := o.Read(&user)
		if err==nil{
			user.Address1 = address1
			_, err1 := o.Update(&user)
			if err1==nil{
				ReturnData["status"] = 1
				ReturnData["message"] = "Address1 changed successfully."
				fmt.Println(user)
			}
		}
	}
	c.Data["json"] = ReturnData
	c.ServeJSON()
}

func (c *MyInformationController) ChangeAddress2() {
	address2 :=c.Input().Get("address2")
	ReturnData := make(map[string]interface{})


	user := models.UserInformation3TG{}
	o := orm.NewOrm()
	o.QueryTable(new(models.UserInformation3TG)).Filter("UserName",c.GetSession("user_name")).One(&user)
	err := o.Read(&user)
	if err==nil{
		user.Address2 = address2
		_, err1 := o.Update(&user)
		if err1==nil{
			ReturnData["status"] = 1
			ReturnData["message"] = "Address2 changed successfully."
			fmt.Println(user)
		}
	}

	c.Data["json"] = ReturnData
	c.ServeJSON()
}

func (c *MyInformationController) ChangeAddress3() {
	address3 :=c.Input().Get("address3")
	ReturnData := make(map[string]interface{})


	user := models.UserInformation3TG{}
	o := orm.NewOrm()
	o.QueryTable(new(models.UserInformation3TG)).Filter("UserName",c.GetSession("user_name")).One(&user)
	err := o.Read(&user)
	if err==nil{
		user.Address3 = address3
		_, err1 := o.Update(&user)
		if err1==nil{
			ReturnData["status"] = 1
			ReturnData["message"] = "Address3 changed successfully."
			fmt.Println(user)
		}
	}

	c.Data["json"] = ReturnData
	c.ServeJSON()
}

