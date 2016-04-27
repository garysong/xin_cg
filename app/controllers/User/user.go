package controllers

//会员管理 登陆
import "fmt"
import "strings"
import "strconv"
import "github.com/revel/revel"
import "xin_cg/app/models"
import "github.com/dchest/captcha"
import "xin_cg/utils"

type User struct {
	*revel.Controller
}

func (c *User) Login(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "登录--深圳幼儿园安全管理"

		CaptchaId := captcha.NewLen(6)

		return c.Render(title, CaptchaId)
	} else {
		var username string = c.Params.Get("username")
		var password string = c.Params.Get("password")

		var captchaId string = c.Params.Get("captchaId")
		var verify string = c.Params.Get("verify")

		data := make(map[string]string)

		if LANG, ok := c.Session["Lang"]; ok {
			//设置语言
			c.Request.Locale = LANG
		} else {
			//设置默认语言
			c.Request.Locale = "zh"
		}

		if !captcha.VerifyString(captchaId, verify) {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("verification_code")
			return c.RenderJson(data)
		}

		if len(username) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_user_name")
			return c.RenderJson(data)
		}

		if len(password) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password")
			return c.RenderJson(data)
		}

		if len(verify) <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_verification_code")
			return c.RenderJson(data)
		}

		admin_info := admin.GetByName(username)

		if admin_info.Id <= 0 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_username_error")
		} else if admin_info.Status == 0 && admin_info.Id != 1 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_forbid_login")
		} else if admin_info.Role.Status == 0 && admin_info.Id != 1 {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("admin_forbid_role_login")
		} else if username == admin_info.Username && utils.Md5(password) == admin_info.Password {
			/*
			 * %% 印出百分比符号，不转换。
			 * %c 整数转成对应的 ASCII 字元。
			 * %d 整数转成十进位。
			 * %f 倍精确度数字转成浮点数。
			 * %o 整数转成八进位。
			 * %s 整数转成字符串。
			 * %x 整数转成小写十六进位。
			 * %X 整数转成大写十六进位
			 */
			c.Session["UserID"] = fmt.Sprintf("%d", admin_info.Id)
			c.Session["Lang"] = admin_info.Lang

			c.Flash.Success(c.Message("login_success"))
			c.Flash.Out["url"] = "/"

			//更新登陆时间
			if ip := c.Request.Header.Get("X-Forwarded-For"); ip != "" {
				ips := strings.Split(ip, ",")
				if len(ips) > 0 && ips[0] != "" {
					rip := strings.Split(ips[0], ":")
					admin.Lastloginip = rip[0]
				}
			} else {
				ip := strings.Split(c.Request.RemoteAddr, ":")
				if len(ip) > 0 {
					if ip[0] != "[" {
						admin.Lastloginip = ip[0]
					}
				}
			}
			admin.UpdateLoginTime(admin_info.Id)

			//管理员日志
			logs := new(models.Logs)
			desc := "登陆用户名:" + admin_info.Username + "|^|登陆系统!|^|登陆ID:" + fmt.Sprintf("%d", admin_info.Id)
			logs.Save(admin_info, c.Controller, desc)

			data["status"] = "1"
			data["url"] = "/Message/"
			data["message"] = c.Message("login_success")
		} else {
			data["status"] = "0"
			data["url"] = "/"
			data["message"] = c.Message("login_password_error")
		}
		return c.RenderJson(data)
	}
}

//退出登陆
func (c *User) Logout(admin *models.Admin) revel.Result {

	if UserID, ok := c.Session["UserID"]; ok {

		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin_info := admin.GetById(UserID)

		//******************************************
		//管理员日志
		logs := new(models.Logs)
		desc := "登陆用户名:" + admin_info.Username + "|^|退出系统!|^|登陆ID:" + fmt.Sprintf("%d", admin_info.Id)
		logs.Save(admin_info, c.Controller, desc)
		//*****************************************

		for k := range c.Session {
			if k != "Lang" {
				delete(c.Session, k)
			}
		}
	}

	c.Flash.Success("安全退出")
	c.Flash.Out["url"] = "/User/Login/"
	return c.Redirect("/Message/")
}

//左侧导航菜单
func (c *User) Left(menu *models.Menu) revel.Result {

	title := "左侧导航--深圳幼儿园安全管理"

	var pid string = c.Params.Get("pid")

	if len(pid) > 0 {
		Pid, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)

			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(Pid, admin_info)

			c.Render(title, left_menu)
		} else {
			c.Render(title)
		}
	} else {
		//获取左侧导航菜单
		//默认获取 首页
		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)

			//获取左侧导航菜单
			left_menu := menu.GetLeftMenuHtml(1, admin_info)

			c.Render(title, left_menu)
		} else {
			c.Render(title)
		}
	}
	return c.RenderTemplate("Public/left.html")
}

//个人信息
func (c *User) EditInfo(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "个人信息--深圳幼儿园安全管理"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)
			c.Render(title, admin_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/EditInfo.html")
	} else {

		var realname string = c.Params.Get("realname")
		if len(realname) > 0 {
			admin.Realname = realname
		} else {
			c.Flash.Error("请输入真实姓名!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		var email string = c.Params.Get("email")
		if len(email) > 0 {
			admin.Email = email
		} else {
			c.Flash.Error("请输入电子邮件!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		var lang string = c.Params.Get("lang")
		if len(lang) > 0 {
			admin.Lang = lang
		} else {
			c.Flash.Error("请选择语言!")
			c.Flash.Out["url"] = "/EditInfo/"
			return c.Redirect("/Message/")
		}

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			if admin.EditInfo(UserID) {

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

//修改密码
func (c *User) EditPwd(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "修改密码--深圳幼儿园安全管理"

		if UserID, ok := c.Session["UserID"]; ok {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)

			c.Render(title, admin_info)
		} else {
			c.Render(title)
		}

		return c.RenderTemplate("User/EditPwd.html")
	} else {

		if UserID, ok := c.Session["UserID"]; ok {

			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
			}

			admin_info := admin.GetById(UserID)

			var old_password string = c.Params.Get("old_password")
			if len(old_password) > 0 {
				if admin_info.Password != utils.Md5(old_password) {
					c.Flash.Error("旧密码不正确!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				}
			} else {
				return c.Redirect("/User/EditPwd/")
			}

			var new_password string = c.Params.Get("new_password")
			if len(new_password) > 0 {

			} else {
				c.Flash.Error("新密码不能为空!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			var new_pwdconfirm string = c.Params.Get("new_pwdconfirm")
			if len(new_pwdconfirm) > 0 {
				if new_pwdconfirm != new_password {
					c.Flash.Error("两次输入密码入不一致!")
					c.Flash.Out["url"] = "/EditPwd/"
					return c.Redirect("/Message/")
				} else {
					admin.Password = new_pwdconfirm
				}
			} else {
				c.Flash.Error("请输入重复新密码!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}

			if admin.EditPwd(UserID) {

				//******************************************
				//管理员日志
				logs := new(models.Logs)
				desc := "个人设置|^|修改密码"
				logs.Save(admin_info, c.Controller, desc)
				//*****************************************

				c.Flash.Success("修改成功!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			} else {
				c.Flash.Error("修改失败!")
				c.Flash.Out["url"] = "/EditPwd/"
				return c.Redirect("/Message/")
			}
		} else {
			c.Flash.Error("未登陆，请先登陆!")
			c.Flash.Out["url"] = "/"
			return c.Redirect("/Message/")
		}
	}
}

//常用菜单
func (c *User) AdminPanel(admin *models.Admin) revel.Result {
	if c.Request.Method == "GET" {
		title := "常用菜单--GoCMS管理系统"

		c.Render(title)

		return c.RenderTemplate("User/AdminPanel.html")
	} else {
		c.Flash.Success("添加成功!")
		c.Flash.Out["url"] = "/AdminPanel/"
		return c.Redirect("/Message/")
	}
}
