package controllers

//import "fmt"

import "github.com/revel/revel"

//验证码
import "github.com/dchest/captcha"

type Captcha struct {
	*revel.Controller
}

//首页
func (c *Captcha) Index() revel.Result {
	var CaptchaId string = c.Params.Get("CaptchaId")
	captcha.WriteImage(c.Response.Out, CaptchaId, 250, 62)
	return nil
}

//返回验证码
func (c *Captcha) GetCaptchaId() revel.Result {
	CaptchaId := captcha.NewLen(6)
	return c.RenderText(CaptchaId)
}
