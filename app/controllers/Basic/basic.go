package controllers

// 管理员设置
//import "strings"
import "strconv"
import "github.com/revel/revel"
import "xin_cg/app/models"
import "xin_cg/app/controllers"

type Basic struct {
	*revel.Controller
}

func (c Basic) Add(basic *models.Basic) revel.Result {

	if c.Request.Method == "GET" {
		title := "添加管理员--深圳幼儿园安全管理"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)

			admin_info := admin.GetById(UserID)

			gardname := controllers.GetNameByCode(admin_info.Code)

			if len(gardname) == 0 {
				c.Render(title)

			} else {
				c.Render(title, gardname)
			}

		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Basic/Add.html")
	} else {
		basic := new(models.Basic)
		var err error
		basic.Code = controllers.GetCodeByName(c.Params.Get("code"))
		basic.Address = c.Params.Get("address")
		revel.WARN.Println(basic.Code)
		revel.WARN.Println(basic.Address)

		basic.Lawpresent = c.Params.Get("lawpresent")
		basic.Lawpresentphone = c.Params.Get("lawpresentphone")
		basic.Director = c.Params.Get("director")
		basic.Directorphone = c.Params.Get("directorphone")
		basic.Contactor = c.Params.Get("contactor")
		basic.Contactorphone = c.Params.Get("contactorphone")
		if Cap := c.Params.Get("capital"); len(Cap) > 0 {
			basic.Capital, err = strconv.ParseFloat(Cap, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		basic.Fax = c.Params.Get("fax")
		basic.Pastalcode = c.Params.Get("pastalcode")
		if floor := c.Params.Get("floorarea"); len(floor) > 0 {
			basic.Floorarea, err = strconv.ParseFloat(floor, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		if built := c.Params.Get("builtarea"); len(built) > 0 {
			basic.Builtarea, err = strconv.ParseFloat(built, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		if outer := c.Params.Get("outerarea"); len(outer) > 0 {
			basic.Outerarea, err = strconv.ParseFloat(outer, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		if staff := c.Params.Get("staffnum"); len(staff) > 0 {
			basic.Staffnum, err = strconv.ParseInt(staff, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		if children := c.Params.Get("childrennum"); len(children) > 0 {
			basic.Childrennum, err = strconv.ParseInt(children, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		if classes := c.Params.Get("classesnum"); len(classes) > 0 {
			basic.Classesnum, err = strconv.ParseInt(classes, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		if security := c.Params.Get("securitynum"); len(security) > 0 {
			basic.Securitynum, err = strconv.ParseInt(security, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}
		}

		basic.Educationattri = c.Params.Get("educationattri")
		basic.Email = c.Params.Get("email")
		basic.Securityinstitution = c.Params.Get("securityinstitution")

		if basic.Save() {

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
				desc := "添加基本信息:" + "|^|基本信息"
				logs.Save(admin_info, c.Controller, desc)
			}

			//*****************************************

			c.Flash.Success("添加基本信息成功!")
			c.Flash.Out["url"] = "/Basic/Info"
			return c.Redirect("/Message/")
		} else {
			c.Flash.Error("添加基本信息失败!")
			c.Flash.Out["url"] = "/Basic/Add/"
			return c.Redirect("/Message/")
		}
	}
}

//个人信息
func (c *Basic) Info(basic *models.Basic) revel.Result {
	if c.Request.Method == "GET" {
		title := "个人信息--深圳幼儿园安全管理"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)

			admin_info := admin.GetById(UserID)
			basic.Code = admin_info.Code

			has := basic.HasBasicInfo()

			gardname := controllers.GetNameByCode(admin_info.Code)

			if has == false {
				c.Render(title, gardname)

				return c.RenderTemplate("Basic/Add.html")
			} else {
				basic_info := basic.GetByCode(admin_info.Code)
				c.Render(title, admin_info, basic_info, gardname)
			}

			c.Render(title, admin_info, gardname)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("Basic/EditInfo.html")
	} else {

		var id string = c.Params.Get("id")
		if len(id) > 0 {
			Id, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			basic.Address = c.Params.Get("address")
			basic.Lawpresent = c.Params.Get("lawpresent")
			basic.Lawpresentphone = c.Params.Get("lawpresentphone")
			basic.Director = c.Params.Get("director")
			basic.Directorphone = c.Params.Get("directorphone")
			revel.WARN.Println(basic.Directorphone)
			basic.Contactor = c.Params.Get("contactor")
			basic.Contactorphone = c.Params.Get("contactorphone")
			if Cap := c.Params.Get("capital"); len(Cap) > 0 {
				basic.Capital, err = strconv.ParseFloat(Cap, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			basic.Fax = c.Params.Get("fax")
			basic.Pastalcode = c.Params.Get("pastalcode")
			if floor := c.Params.Get("floorarea"); len(floor) > 0 {
				basic.Floorarea, err = strconv.ParseFloat(floor, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			if built := c.Params.Get("builtarea"); len(built) > 0 {
				basic.Builtarea, err = strconv.ParseFloat(built, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			if outer := c.Params.Get("outerarea"); len(outer) > 0 {
				basic.Outerarea, err = strconv.ParseFloat(outer, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			if staff := c.Params.Get("staffnum"); len(staff) > 0 {
				basic.Staffnum, err = strconv.ParseInt(staff, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			if children := c.Params.Get("childrennum"); len(children) > 0 {
				basic.Childrennum, err = strconv.ParseInt(children, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			if classes := c.Params.Get("classesnum"); len(classes) > 0 {
				basic.Classesnum, err = strconv.ParseInt(classes, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			if security := c.Params.Get("securitynum"); len(security) > 0 {
				basic.Securitynum, err = strconv.ParseInt(security, 10, 64)
				if err != nil {
					revel.WARN.Println(err)
				}
			}

			basic.Educationattri = c.Params.Get("educationattri")
			basic.Email = c.Params.Get("email")
			basic.Securityinstitution = c.Params.Get("securityinstitution")

			revel.WARN.Println(Id)
			if basic.Edit(Id) {
				if UserID, ok := c.Session["UserID"]; ok {
					UserID, err := strconv.ParseInt(UserID, 10, 64)
					if err != nil {
						revel.WARN.Println(err)
					}

					admin := new(models.Admin)
					admin_info := admin.GetById(UserID)

					c.Session["Lang"] = admin_info.Lang

					logs := new(models.Logs)
					desc := "个人设置|^|个人信息"
					logs.Save(admin_info, c.Controller, desc)

				}
				c.Flash.Success(c.Message("operation_success"))
				c.Flash.Out["url"] = "/BasicInfo/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error(c.Message("operation_failure"))
				c.Flash.Out["url"] = "/BasicInfo/"
				return c.Redirect("/Message/")
			}

		} else {
			c.Flash.Error("编辑基本信息失败!")
			c.Flash.Out["url"] = "/Basic/Info/"
			return c.Redirect("/Message/")
		}

	}
}
