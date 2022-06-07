package routers

import (
	"3TrashGuitarists/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/about_us", &controllers.AboutUsController{})
	beego.Router("/charvel", &controllers.CharvelController{})
	beego.Router("/esp", &controllers.ESPController{})
	beego.Router("/fender", &controllers.FenderController{})
	beego.Router("/ibanez", &controllers.IbanezController{})
	beego.Router("/suhr", &controllers.SuhrController{})
	beego.Router("/tom_anderson", &controllers.TAController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/sign_in", &controllers.SignInController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/myinformation", &controllers.MyInformationController{})
	beego.Router("/myinformation_changepassword", &controllers.MyInformationController{}, "post:ChangePassword")
	beego.Router("/myinformation_changeaddress1", &controllers.MyInformationController{}, "post:ChangeAddress1")
	beego.Router("/myinformation_changeaddress2", &controllers.MyInformationController{}, "post:ChangeAddress2")
	beego.Router("/myinformation_changeaddress3", &controllers.MyInformationController{}, "post:ChangeAddress3")
	//beego.Router("/GuitarDisplay", &controllers.GuitarDisplayController{})
	//beego.Router("/guitar_display/:brand/:series", &controllers.GuitarDisplayController{})
	beego.Router("/guitar_display/:brand/:series/:neck/:fretboard/:body", &controllers.GuitarDisplayController{})
	beego.Router("/shop_details/:id", &controllers.ShopDetailsController{})
	beego.Router("/checkout/:id/:amount", &controllers.CheckoutController{})
	beego.Router("/orders", &controllers.OrdersController{})

}
