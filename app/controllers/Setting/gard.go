package controllers

// 管理员设置
import "strconv"
import "github.com/revel/revel"
import "xin_cg/app/models"

type Gard struct {
	*revel.Controller
}

//首页
func (c Gard) Index(gard *models.Gard) revel.Result {
	title := "管理员管理--深圳幼儿园安全管理"

	var page string = c.Params.Get("page")

	if len(page) > 0 {

		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {

			revel.WARN.Println(err)
		}

		gard_list, pages := gard.GetByAll(Page, 10)

		c.Render(title, gard_list, pages)
	} else {

		gard_list, pages := gard.GetByAll(1, 10)

		c.Render(title, gard_list, pages)
	}

	return c.RenderTemplate("Setting/Gard/Index.html")
}

//添加管理员
func (c Gard) Add(gard *models.Gard) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加管理员--深圳幼儿园安全管理"

		g := new(models.Gard)
		gard_list := g.GetDictList()

		c.Render(title, gard_list)
		return c.RenderTemplate("Setting/Gard/Add.html")
	} else {

		var dict string = c.Params.Get("dict")
		if len(dict) > 0 {
			gard.Dict = dict
		} else {
			c.Flash.Error("请输入区名称!")
			c.Flash.Out["url"] = "/Gard/Add/"
			return c.Redirect("/Message/")
		}

		var street string = c.Params.Get("Street")
		if len(street) > 0 {
			gard.Street = street
		}

		var childrengard string = c.Params.Get("Childrengard")
		if len(childrengard) > 0 {
			gard.Childrengard = childrengard
		}

		var code string = c.Params.Get("Code")
		if len(code) > 0 {
			gard.Code = code
		}

		if len(gard.Code) > 0 && gard.HasGard() {
			c.Flash.Error("幼儿园编号“" + gard.Code + "”已存在！")
			c.Flash.Out["url"] = "/Gard/Add/"
			return c.Redirect("/Message/")
		}

		if gard.Save() {

			//******************************************
			//管理员日志
			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				desc := "添加幼儿园:" + gard.Code + "|^|幼儿园管理"
				logs.Save(admin_info, c.Controller, desc)
			}

			//*****************************************

			c.Flash.Success("添加幼儿园成功!")
			c.Flash.Out["url"] = "/Gard/"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加幼儿园失败!")
			c.Flash.Out["url"] = "/Gard/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//编辑幼儿园
func (c Gard) Edit(gard *models.Gard) revel.Result {
	if c.Request.Method == "GET" {
		title := "编辑管理员--深圳幼儿园安全管理"

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			g := new(models.Gard)

			gard_list := g.GetDictList()

			gard_info := gard.GetById(Id)

			c.Render(title, gard_info, gard_list)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Setting/Gard/Edit.html")
	} else {

		var id string = c.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			var dict string = c.Params.Get("dict")

			if len(dict) > 0 {
				gard.Dict = dict
			} else {
				c.Flash.Error("请输入区名称!")
				c.Flash.Out["url"] = "/Gard/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}

			var street string = c.Params.Get("Street")
			if len(street) > 0 {
				gard.Street = street
			}

			var childrengard string = c.Params.Get("Childrengard")
			if len(childrengard) > 0 {
				gard.Childrengard = childrengard
			}

			var code string = c.Params.Get("Code")
			if len(code) > 0 {
				gard.Code = code
			}

			if len(gard.Code) > 0 && gard.HasGard() {
				c.Flash.Error("幼儿园编号“" + gard.Code + "”已存在！")
				c.Flash.Out["url"] = "/Gard/Add/"
				return c.Redirect("/Message/")
			}

			if gard.Edit(Id) {

				//******************************************
				//管理员日志
				if UserID, ok := c.Session["UserID"]; ok {
					UserID, err := strconv.ParseInt(UserID, 10, 64)
					if err != nil {
						revel.WARN.Println(err)
					}

					admin := new(models.Admin)
					admin_info := admin.GetById(UserID)

					logs := new(models.Logs)
					desc := "编辑幼儿园:" + gard.Code + "|^|幼儿园管理"
					logs.Save(admin_info, c.Controller, desc)
				}
				//*****************************************

				c.Flash.Success("编辑幼儿园成功!")
				c.Flash.Out["url"] = "/Gard/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("编辑幼儿园失败!")
				c.Flash.Out["url"] = "/Gard/Edit/" + id + "/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("编辑幼儿园失败!")
			c.Flash.Out["url"] = "/Gard/Edit/" + id + "/"
			return c.Redirect("/Message/")
		}

	}
}

//删除幼儿园
func (c Gard) Delete(gard *models.Gard) revel.Result {
	var id string = c.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if gard.DelByID(Id) {

			//******************************************
			//管理员日志
			if UserID, ok := c.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)
				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				desc := "删除幼儿园|^|ID:" + id
				logs.Save(admin_info, c.Controller, desc)
			}
			//*****************************************

			data["status"] = "1"
			data["message"] = "删除成功!"
			return c.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "删除失败!"
			return c.RenderJson(data)
		}
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
		return c.RenderJson(data)
	}
}
