package controllers

import (
	"3TrashGuitarists/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ShopDetailsController struct {
	//beego.Controller
	BaseController
}

func (c *ShopDetailsController) Get() {
	id := c.Ctx.Input.Param(":id")

	o := orm.NewOrm()
	product := models.Guitars{}
	o.QueryTable(new(models.Guitars)).Filter("Id",id).One(&product)

	c.Data["product"] = product
	fmt.Println(id)
	fmt.Println(c.Data["product"])


	c.TplName = "ShopDetails.html"

}
