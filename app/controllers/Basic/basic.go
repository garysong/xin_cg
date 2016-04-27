package controllers

// 管理员设置
import "strconv"
import "github.com/revel/revel"
import "xin_cg/app/models"

type Basic struct {
	*revel.Controller
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

			basic_info := basic.GetByCode(admin_info.Code)

			if basic_info == nil {
				c.Render(title, admin_info)
				return c.RenderTemplate("Basic/EditInfo.html")
			} else {
				c.Render(title, admin_info, basic_info)
			}

			c.Render(title, basic_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/EditInfo.html")
	} else {

		basic.Lawpresent = c.Params.Get("lawpresent")
		basic.Lawpresentphone = c.Params.Get("lawpresentphone")
		basic.Director = c.Params.Get("director")
		basic.Directorphone = c.Params.Get("directorphone")
		basic.Contactor = c.Params.Get("contactor")
		basic.Contactorphone = c.Params.Get("contactorphone")
		basic.Capital = c.Params.Get("capital")
		basic.Fax = c.Params.Get("fax")
		basic.Pastalcode = c.Params.Get("pastalcode")
		basic.Floorarea = c.Params.Get("floorarea")
		basic.Builtarea = c.Params.Get("builtarea")
		basic.Outerarea = c.Params.Get("outerarea")
		basic.Staffnum = c.Params.Get("staffnum")
		basic.Childrennum = c.Params.Get("childrennum")
		basic.Classesnum = c.Params.Get("classesnum")
		basic.Securitynum = c.Params.Get("securitynum")
		basic.Educationattri = c.Params.Get("educationattri")
		basic.Email = c.Params.Get("Email")
		basic.Securityinstitution = c.Params.Get("secrityinstitution")

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			if basic.Edit(UserID) {

				//******************************************
				//管理员日志
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

				if LANG, ok := c.Session["Lang"]; ok {
					//设置语言
					c.Request.Locale = LANG
				} else {
					//设置默认语言
					c.Request.Locale = "zh"
				}

				c.Flash.Success(c.Message("operation_success"))
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error(c.Message("operation_failure"))
				c.Flash.Out["url"] = "/EditInfo/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error(c.Message("not_login"))
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}
