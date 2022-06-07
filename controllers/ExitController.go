package controllers

type ExitController struct {
	BaseController
}

func (c *ExitController)Get(){
	c.DelSession("user_name")
	c.Redirect("/",302)
}

