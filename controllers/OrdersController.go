package controllers

import (
	"3TrashGuitarists/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type OrdersController struct {
	//beego.Controller
	BaseController
}

type OrderInformation struct {
	OrderNumber string
	ProductName string
	Price float64
	Quantity int
	Total float64
	Status string
	GuitarID string
}


func (c *OrdersController) Get() {
	user_name := c.GetSession("user_name")
	fmt.Println("user_name---->", user_name)
	if user_name != nil {
		c.IsLogin = true
		c.user_name = user_name
	} else {
		c.IsLogin = false
		c.Redirect(beego.URLFor("/sign_in"),302)
	}

	o := orm.NewOrm()
	user := models.UserInformation3TG{}
	o.QueryTable(new(models.UserInformation3TG)).Filter("user_name",user_name).One(&user)
	orders := []models.Order{}
	o.QueryTable(new(models.Order)).Filter("UserInformation3TG",&user).All(&orders)
	fmt.Println(orders[1])

	var orderInformation []OrderInformation
	for _,v := range orders{
		guitar := models.Guitars{}
		o.QueryTable(new(models.Guitars)).Filter("Id",*v.Guitars).One(&guitar)

		order := OrderInformation{
			OrderNumber: strconv.Itoa(v.Id),
			ProductName: guitar.GuitarName,
			Price: guitar.Price,
			Quantity: v.Amount,
			Total: v.Total,
			Status: "Preparing to ship",
			GuitarID: strconv.Itoa(guitar.Id)}
		//order :=new(OrderInformation)
		//order.orderNumber = string(v.Id)
		//order.productName = v.Guitars.GuitarName
		//order.price = v.Guitars.Price
		//order.quantity = v.Amount
		//order.total = v.Total
		//order.status = "Preparing to ship"
		orderInformation = append(orderInformation, order)
	}

	c.Data["orderInformation"] = orderInformation
	fmt.Println(c.Data["orderInformation"])
	c.TplName = "Orders.html"

}


func (c *OrdersController) Post() {
	firstName :=c.Input().Get("firstName")
	lastName :=c.Input().Get("lastName")
	phoneNumber :=c.Input().Get("phoneNumber")
	companyName :=c.Input().Get("companyName")
	orderNotes :=c.Input().Get("orderNotes")
	address1Checked :=c.Input().Get("address1Checked")
	address2Checked :=c.Input().Get("address2Checked")
	//address3Checked :=c.Input().Get("address3Checked")
	bank :=c.Input().Get("bank")
	cash :=c.Input().Get("cash")
	//paypal :=c.Input().Get("paypal")
	userId,_ :=strconv.Atoi(c.Input().Get("userId"))
	guitarId,_ :=strconv.Atoi(c.Input().Get("guitarId"))
	amount,_ := strconv.Atoi(c.Input().Get("amount"))
	total,_:= strconv.ParseFloat(c.Input().Get("total"), 64)




	ReturnData := make(map[string]interface{})
	o := orm.NewOrm()
	guitar := models.Guitars{}
	o.QueryTable(new(models.Guitars)).Filter("Id",guitarId).One(&guitar)
	if amount > guitar.StockNumber{
		ReturnData["status"]= 0
		ReturnData["message"]="The quantity of this product is insufficient."
	}else{
		order :=models.Order{}
		user := models.UserInformation3TG{}
		o.QueryTable(new(models.UserInformation3TG)).Filter("Id",userId).One(&user)
		order.UserInformation3TG = &user
		order.Guitars = &guitar
		order.Amount = amount
		order.Total = total
		order.FirstName = firstName
		order.LastName = lastName
		order.Phone = phoneNumber
		order.CompanyName = companyName
		order.Notes = orderNotes
		if address1Checked != ""{
			order.Address = order.UserInformation3TG.Address1
		}else if address2Checked != ""{
			order.Address = order.UserInformation3TG.Address2
		}else{
			order.Address = order.UserInformation3TG.Address3
		}

		if bank != ""{
			order.PaymentMethod = 0
		}else if cash != ""{
			order.PaymentMethod = 1
		}else{
			order.PaymentMethod = -1
		}

		id, err := o.Insert(&order)
		if err != nil {
			ReturnData["status"] = -1
			ReturnData["message"] = "fail to insert order into SQL."
			fmt.Println("fail to insert order into SQL. ERR: %v",err)
		} else {
			fmt.Println("the id for this order is: ", id)
			guitar.StockNumber = guitar.StockNumber - amount
			_, err1 := o.Update(&guitar)
			if err1==nil{
				ReturnData["status"] = 1
				ReturnData["message"] = "Order successful."
				fmt.Println(guitar.StockNumber)
			}
		}
	}

	fmt.Println(ReturnData)
	c.Data["json"] = ReturnData
	c.ServeJSON()
}