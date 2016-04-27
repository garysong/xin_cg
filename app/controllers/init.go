package controllers

//初始化入口文件
import "runtime"
import "strconv"
import "path/filepath"
import "xin_cg/utils"
import "github.com/revel/revel"
import "github.com/revel/revel/modules/jobs/app/jobs"
import "xin_cg/app/models"

type Name struct {
	Dict         string
	Street       string
	Childrengard string
}

var BasePath, _ = filepath.Abs("")
var Code2Name = utils.NewBeeMap()

//定义项目根目录
var ROOT_DIR string = BasePath

//定义项目上传文件目录
var UPLOAD_DIR string = BasePath + "/www/upload/"

func init() {
	revel.OnAppStart(Bootstrap)

	//检测是否登陆
	revel.InterceptFunc(CheckLogin, revel.BEFORE, revel.ALL_CONTROLLERS)
}

//系统初始化变量
func Bootstrap() {
	//多核运行
	np := runtime.NumCPU()
	if np >= 2 {
		runtime.GOMAXPROCS(np - 1)
	}

	if runtime.GOOS == "windows" {
		UPLOAD_DIR = BasePath + "\\www\\upload\\"
	} else {
		UPLOAD_DIR = BasePath + "/www/upload/"
	}

	GetAllCGCode()

	//任务调度模块

	//前6个字段分别表示：
	//       秒钟：0-59
	//       分钟：0-59
	//       小时：1-23
	//       日期：1-31
	//       月份：1-12
	//       星期：0-6（0 表示周日）

	//还可以用一些特殊符号：
	//       *： 表示任何时刻
	//       ,：　表示分割，如第三段里：2,4，表示 2 点和 4 点执行
	//　　    －：表示一个段，如第三端里： 1-5，就表示 1 到 5 点
	//       /n : 表示每个n的单位执行一次，如第三段里，*/1, 就表示每隔 1 个小时执行一次命令。也可以写成1-23/1.
	/////////////////////////////////////////////////////////
	//  0/30 * * * * *                        每 30 秒 执行
	//  0 43 21 * * *                         21:43 执行
	//  0 15 05 * * * 　　                     05:15 执行
	//  0 0 17 * * *                          17:00 执行
	//  0 0 17 * * 1                          每周一的 17:00 执行
	//  0 0,10 17 * * 0,2,3                   每周日,周二,周三的 17:00和 17:10 执行
	//  0 0-10 17 1 * *                       毎月1日从 17:00 到 7:10 毎隔 1 分钟 执行
	//  0 0 0 1,15 * 1                        毎月1日和 15 日和 一日的 0:00 执行
	//  0 42 4 1 * * 　 　                     毎月1日的 4:42 分 执行
	//  0 0 21 * * 1-6　　                     周一到周六 21:00 执行
	//  0 0,10,20,30,40,50 * * * *　           每隔 10 分 执行
	//  0 */10 * * * * 　　　　　　              每隔 10 分 执行
	//  0 * 1 * * *　　　　　　　　               从 1:0 到 1:59 每隔 1 分钟 执行
	//  0 0 1 * * *　　　　　　　　               1:00 执行
	//  0 0 */1 * * *　　　　　　　               毎时 0 分 每隔 1 小时 执行
	//  0 0 * * * *　　　　　　　　               毎时 0 分 每隔 1 小时 执行
	//  0 2 8-20/3 * * *　　　　　　             8:02,11:02,14:02,17:02,20:02 执行
	//  0 30 5 1,15 * *　　　　　　              1 日 和 15 日的 5:30 执行

	//每分钟执行         0 * * * * *
	jobs.Schedule("0 * * * * *", models.EveryMinute{})

	//每五分钟执行    0 5 * * * *
	jobs.Schedule("0 5 * * * *", models.FiveMinutes{})

	//每三十分钟执行    0 30 * * * *
	jobs.Schedule("0 30 * * * *", models.ThirtyMinutes{})

	//每小时执行        0 0 * * * *
	jobs.Schedule("0 0 * * * *", models.HourlyMinutes{})

	//每天执行          @daily
	//每天运行一次,午夜
	jobs.Schedule("@daily", models.DailyMinutes{})

	//每周执行          @weekly
	//每周运行一次,周日午夜
	jobs.Schedule("@weekly", models.WeeklyMinutes{})

	//每月执行          @monthly
	//一个月运行一次,半夜,第一个月
	jobs.Schedule("@monthly", models.MonthlyMinutes{})

	//每年执行          @yearly
	//运行一年一次,1月1日午夜
	jobs.Schedule("@yearly", models.YearlyMinutes{})
}

//检测登陆
func CheckLogin(c *revel.Controller) revel.Result {

	//登陆页面，CSS, JS, Ajax, 验证码页面 都不进行登陆验证
	if c.Name == "User" && c.MethodName == "Login" || c.Name == "Ajax" || c.Name == "Static" || c.Name == "Captcha" || c.Name == "Kindeditor" {

		if LANG, ok := c.Session["Lang"]; ok {
			//设置语言
			c.RenderArgs["currentLocale"] = LANG
		} else {
			//设置默认语言
			c.RenderArgs["currentLocale"] = "zh"
		}

		return nil
	} else {

		UserID := utils.GetSession("UserID", c.Session)

		if len(UserID) > 0 {
			UserID, err := strconv.ParseInt(UserID, 10, 64)
			if err != nil {
				revel.WARN.Println(err)
				return c.Redirect("/Login/")
			}

			admin := new(models.Admin)
			admin_info := admin.GetById(UserID)
			if admin_info.Id <= 0 {
				return c.Redirect("/Login/")
			}

			//控制器
			c.RenderArgs["Controller"] = c.Name
			//动作
			c.RenderArgs["action"] = c.Action
			//模型
			c.RenderArgs["Model"] = c.MethodName

			//登陆信息
			c.RenderArgs["admin_info"] = admin_info

			//设置语言
			c.RenderArgs["currentLocale"] = admin_info.Lang
		} else {

			//控制器
			c.RenderArgs["Controller"] = c.Name
			//动作
			c.RenderArgs["action"] = c.Action
			//模型
			c.RenderArgs["Model"] = c.MethodName

			return c.Redirect("/Login/")
		}
	}

	return nil
}

func GetAllCGCode() {
	gard := new(models.Gard)
	gard_list, _ := gard.GetByAll(0, 10000)

	revel.WARN.Println(gard_list)
	for _, g := range gard_list {
		n := new(Name)
		n.Dict = g.Dict
		n.Street = g.Street
		n.Childrengard = g.Childrengard
		Code2Name.Set(g.Code, n)
	}
}
