package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserInformation3TG struct {
	Id int `orm:"pk;auto"`
	UserName string `orm:"index;unique"`
	Password string
	Email string `orm:"index;unique"`
	CreatTime time.Time `orm:"auto_now_add;type(datetime)"`
	MoneyRemain float64
	Address1 string
	Address2 string
	Address3 string
	Order []*Order `orm:"reverse(many)"`
}

type Guitars struct {
	Id int `orm:"pk;auto"`
	GuitarName string
	ModelNumber string
	BrandName string
	SeriesName string
	NeckWood string
	FingerBoardWood string
	BodyWood string
	FretsNumber int
	TopWood string
	DisplayCover string
	DisplayPicture1 string
	DisplayPicture2 string
	DisplayPicture3 string
	DisplayPicture4 string
	DisplayPicture5 string
	Description string `orm:"size(10000)"`
	Price float64
	StockNumber int `orm:"default(0)"`
	Order []*Order `orm:"reverse(many)"`
}

type Order struct{
	Id int `orm:"pk;auto"`
	UserInformation3TG *UserInformation3TG `orm:"rel(fk)"`
	Guitars *Guitars `orm:"rel(fk)"`
	Amount int
	Total float64
	FirstName string
	LastName string
	Phone string
	CompanyName string
	Address string
	Notes string
	PaymentMethod int // 0 for bank transfer, 1 for cash, -1 for paypal
}



func (u *UserInformation3TG)TableName() string{
	return "user_information"
}

func (g *Guitars)TableName() string{
	return "guitars"
}


func (o *Order)TableName() string{
	return "order"
}

func init() {
	orm.RegisterModel(new(UserInformation3TG))
	orm.RegisterModel(new(Guitars))
	orm.RegisterModel(new(Order))
}