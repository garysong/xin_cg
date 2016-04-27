package controllers

//我的面板首页
import "github.com/revel/revel"

type Panel struct {
	*revel.Controller
}

func (c Panel) Index() revel.Result {
	title := "设置--深圳幼儿园安全管理"

	c.Render(title)
	return c.RenderTemplate("Panel/Index.html")
}
