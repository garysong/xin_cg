package controllers

//基本情况--消防设施
import "strconv"
import "github.com/revel/revel"
import "xin_cg/app/models"

//import "xin_cg/utils"
import "xin_cg/app/controllers"

type Equipment struct {
	*revel.Controller
}

func (e Equipment) Index(equipment *models.Equipment) revel.Result {
	title := "消防设施--深圳幼儿园安全管理"

	var page string = e.Params.Get("page")

	where := make(map[string]string)

	code := controllers.GetCode(e.Controller)

	gardname := controllers.GetNameByCode(code)

	if len(page) > 0 {
		Page, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		equipment_list, pages := equipment.GetByCode(code, where, Page, 10)

		e.Render(title, equipment_list, pages, gardname)
	} else {
		equipment_list, pages := equipment.GetByCode(code, where, 1, 10)
		e.Render(title, equipment_list, pages, gardname)
	}

	return e.RenderTemplate("Basic/Equipment/Index.html")
}

//添加消防设施
func (e Equipment) Add(equipment *models.Equipment) revel.Result {
	if e.Request.Method == "GET" {
		title := "添加消防设施--深圳幼儿园安全管理"

		code := controllers.GetCode(e.Controller)

		gardname := controllers.GetNameByCode(code)

		if len(gardname) == 0 {
			e.Render(title)

		} else {
			e.Render(title, gardname)
		}
		return e.RenderTemplate("Basic/Equipment/Add.html")
	} else {

		var err error

		equipment.Code = controllers.GetCodeByName(e.Params.Get("code"))
		equipment.Name = e.Params.Get("name")
		equipment.Model = e.Params.Get("model")

		if num := e.Params.Get("num"); len(num) > 0 {
			equipment.Num, err = strconv.ParseInt(num, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}
		equipment.Status = e.Params.Get("status")

		if equipment.Save() {

			//******************************************
			//管理员日志
			if UserID, ok := e.Session["UserID"]; ok {
				UserID, err := strconv.ParseInt(UserID, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}

				admin := new(models.Admin)

				admin_info := admin.GetById(UserID)

				logs := new(models.Logs)
				desc := "添加消防设施:" + "|^|基本信息"
				logs.Save(admin_info, e.Controller, desc)
			}

			//*****************************************

			e.Flash.Success("添加基本信息成功!")
			e.Flash.Out["url"] = "/Equipment"
			return e.Redirect("/Message/")
		} else {
			e.Flash.Error("添加基本信息失败!")
			e.Flash.Out["url"] = "/Equipment"
			return e.Redirect("/Message/")
		}

	}
}

func (e Equipment) Edit(equipment *models.Equipment) revel.Result {

	if e.Request.Method == "GET" {
		title := "消防设施--深圳幼儿园安全管理"

		var id string = e.Params.Get("id")

		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			code := controllers.GetCode(e.Controller)

			if len(code) > 0 {
				gardname := controllers.GetNameByCode(code)
				equipment_info := equipment.GetById(Id)

				e.Render(title, equipment_info, gardname)

			} else {
				e.Render(title)
			}
		} else {
			e.Render(title)
		}

		return e.RenderTemplate("Basic/Equipment/Edit.html")

	} else {

		var id string = e.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)

			if err != nil {
				revel.WARN.Println(err)
			}

			equipment.Code = controllers.GetCodeByName(e.Params.Get("code"))
			equipment.Name = e.Params.Get("name")
			equipment.Model = e.Params.Get("model")

			if num := e.Params.Get("num"); len(num) > 0 {
				equipment.Num, err = strconv.ParseInt(num, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}
			equipment.Status = e.Params.Get("status")

			if equipment.Edit(Id) {
				if UserID, ok := e.Session["UserID"]; ok {
					UserID, err := strconv.ParseInt(UserID, 10, 64)
					if err != nil {
						revel.WARN.Println(err)
					}

					admin := new(models.Admin)
					admin_info := admin.GetById(UserID)

					e.Session["Lang"] = admin_info.Lang

					logs := new(models.Logs)
					desc := "编辑消防设施：" + "|^|基本信息^|ID:" + id
					logs.Save(admin_info, e.Controller, desc)

				}
				e.Flash.Success("编辑消防设施成功")
				e.Flash.Out["url"] = "/Equipment"
				return e.Redirect("/Message/")
			} else {
				e.Flash.Error("编辑消防设施失败")
				e.Flash.Out["url"] = "/Equipment/Edit/" + id + "/"
				return e.Redirect("/Message/")
			}

		} else {
			e.Flash.Error("编辑基本信息失败!")
			e.Flash.Out["url"] = "/Equipment/Edit/" + id + "/"
			return e.Redirect("/Message/")
		}

	}
}

//删除消防设施
func (e Equipment) Delete(equipment *models.Equipment) revel.Result {

	var id string = e.Params.Get("id")

	data := make(map[string]string)

	if len(id) > 0 {
		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if equipment.DelByID(Id) {
			data["status"] = "1"
			data["message"] = "删除成功!"
			return e.RenderJson(data)
		} else {
			data["status"] = "0"
			data["message"] = "删除失败!"
			return e.RenderJson(data)
		}
	} else {
		data["status"] = "0"
		data["message"] = "删除失败!"
		return e.RenderJson(data)
	}
}
