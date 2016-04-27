package controllers

//后台公用处理
//import "strconv"
import "github.com/revel/revel"

//import "admin/app/models"

type Public struct {
	*revel.Controller
}

//消息提示
func (c *Public) Message() revel.Result {
	c.Render()
	return c.RenderTemplate("Public/message.html")
}
