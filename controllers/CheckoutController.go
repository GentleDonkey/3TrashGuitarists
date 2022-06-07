package controllers

import (
	"3TrashGuitarists/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
)

type CheckoutController struct {
	//beego.Controller
	BaseController
}

func (c *CheckoutController) Get() {
	user_name := c.GetSession("user_name")
	fmt.Println("user_name---->", user_name)
	if user_name != nil {
		c.IsLogin = true
		c.user_name = user_name
	} else {
		c.IsLogin = false
		c.Redirect(beego.URLFor("/sign_in"),302)
	}

	id := c.Ctx.Input.Param(":id")
	amount_string := c.Ctx.Input.Param(":amount")
	amount, err := strconv.ParseFloat(amount_string, 64)
	if err == nil{
		fmt.Println(id,amount+amount)
	} else {
		fmt.Println("error in type transformation from string to int")
	}

	o := orm.NewOrm()
	guitar := models.Guitars{}
	o.QueryTable(new(models.Guitars)).Filter("Id",id).One(&guitar)

	user := models.UserInformation3TG{}
	o.QueryTable(new(models.UserInformation3TG)).Filter("UserName",user_name).One(&user)

	user.Password = ""

	var shipping_fee float64 =50
	tax := 0.13
	var total = (guitar.Price * amount + shipping_fee) * (tax+1)
	total = math.Trunc(total*100+0.5)*0.01
	c.Data["guitar"]=guitar
	c.Data["user"]=user
	c.Data["shipping_fee"]=shipping_fee
	c.Data["amount"]=amount
	c.Data["tax"]=fmt.Sprintf("%.2f", tax*(guitar.Price * amount + shipping_fee))
	c.Data["total"]=fmt.Sprintf("%.2f", total)
	fmt.Println(c.Data["guitar"],c.Data["user"],c.Data["amount"])

	c.TplName = "Checkout.html"

}
