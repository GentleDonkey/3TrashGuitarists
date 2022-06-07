package controllers

import (
	"3TrashGuitarists/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GuitarDisplayController struct {
	beego.Controller
}

func (c *GuitarDisplayController) Get() {
	brand := c.Ctx.Input.Param(":brand")
	series := c.Ctx.Input.Param(":series")
	neck := c.Ctx.Input.Param(":neck")
	fretboard := c.Ctx.Input.Param(":fretboard")
	body := c.Ctx.Input.Param(":body")

	o := orm.NewOrm()
	guitars := []models.Guitars{}
	if brand == "all" {
		if series == "all" {
			if neck == "all" {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			} else {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			}
		} else {
			if neck == "all" {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			} else {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("NeckWood", neck).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("NeckWood", neck).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("SeriesName", series).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			}
		}
	} else {
		if series == "all" {
			if neck == "all" {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			} else {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			}
		} else {
			if neck == "all" {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			} else {
				if fretboard == "all" {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("NeckWood", neck).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("NeckWood", neck).Filter("BodyWood", body).All(&guitars)
					}
				} else {
					if body == "all" {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).All(&guitars)
					} else {
						o.QueryTable(new(models.Guitars)).Filter("BrandName", brand).Filter("SeriesName", series).Filter("NeckWood", neck).Filter("FingerBoardWood", fretboard).Filter("BodyWood", body).All(&guitars)
					}
				}
			}
		}
	}
	c.Data["guitars"] = guitars
	fmt.Println(c.Data["guitars"])
	c.TplName = "GuitarDisplay.html"

}
