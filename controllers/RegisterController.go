package controllers

import (
	"3TrashGuitarists/models"
	"3TrashGuitarists/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "Register.html"

}

func (c *RegisterController) Post() {
	ReturnData := make(map[string]interface{})

	username :=c.Input().Get("username")
	password :=c.Input().Get("password")
	repassword :=c.Input().Get("repassword")
	email :=c.Input().Get("email")
	address1 :=c.Input().Get("address1")

	ReturnData["status"] =1 //1: sign up successfully, 0: wrong format,-1: fail to insert user data into sql
	ReturnData["message"] ="Thank you for joining us!"
	ReturnData["check_username_dup"] =false
	ReturnData["check_username_format"] =false
	ReturnData["check_password"] =false
	ReturnData["check_repassword"] =false
	ReturnData["check_email"] =false
	ReturnData["check_email_dup"] =false
	ReturnData["check_address1"] =false


	o := orm.NewOrm()
	user_exist := o.QueryTable(new(models.UserInformation3TG)).Filter("user_name",username).Exist()
	email_exist := o.QueryTable(new(models.UserInformation3TG)).Filter("email",email).Exist()


	if user_exist == true {
		ReturnData["check_username_dup"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}
	if utils.VerifyNameFormat(username) == false {
		ReturnData["check_username_format"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}
	if utils.VerifyPasswordFormat(password) == false {
		ReturnData["check_password"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}
	if repassword != password {
		ReturnData["check_repassword"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}
	if utils.VerifyEmailFormat(email) == false {
		ReturnData["check_email"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}
	if email_exist == true {
		ReturnData["check_email_dup"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}
	if address1 == "" {
		ReturnData["check_address1"] = true
		ReturnData["status"] = 0
		ReturnData["message"] = ""
	}

	if ReturnData["status"] == 1 {
		user_information := models.UserInformation3TG{}
		user_information.UserName = username
		user_information.Password = utils.GetMd5(password)
		user_information.Email = email
		user_information.Address1 = address1
		user_information.MoneyRemain =0
		user_information.CreatTime = time.Now()

		id, err := o.Insert(&user_information)
		if err != nil {
			ReturnData["status"] = -1
			ReturnData["message"] = "fail to insert user data into SQL."
			fmt.Println("fail to insert user data into SQL. ERR: %v",err)
		} else {
			fmt.Println("the id for this user is: ", id)
			c.SetSession("user_name",username)
		}
	}

	fmt.Println(ReturnData)
	c.Data["json"] = ReturnData
	c.ServeJSON()
}
