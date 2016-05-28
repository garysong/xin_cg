// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Main(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.Main", args).Url
}


type tBasic struct {}
var Basic tBasic


func (_ tBasic) Add(
		basic interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "basic", basic)
	return revel.MainRouter.Reverse("Basic.Add", args).Url
}

func (_ tBasic) Info(
		basic interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "basic", basic)
	return revel.MainRouter.Reverse("Basic.Info", args).Url
}


type tEquipment struct {}
var Equipment tEquipment


func (_ tEquipment) Index(
		equipment interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "equipment", equipment)
	return revel.MainRouter.Reverse("Equipment.Index", args).Url
}

func (_ tEquipment) Add(
		equipment interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "equipment", equipment)
	return revel.MainRouter.Reverse("Equipment.Add", args).Url
}

func (_ tEquipment) Edit(
		equipment interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "equipment", equipment)
	return revel.MainRouter.Reverse("Equipment.Edit", args).Url
}

func (_ tEquipment) Delete(
		equipment interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "equipment", equipment)
	return revel.MainRouter.Reverse("Equipment.Delete", args).Url
}


type tPanel struct {}
var Panel tPanel


func (_ tPanel) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Panel.Index", args).Url
}


type tAjax struct {}
var Ajax tAjax


func (_ tAjax) GetCaptcha(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetCaptcha", args).Url
}

func (_ tAjax) Pos(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Ajax.Pos", args).Url
}

func (_ tAjax) GetPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.GetPanel", args).Url
}

func (_ tAjax) DelPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.DelPanel", args).Url
}

func (_ tAjax) AddPanel(
		admin_panel interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin_panel", admin_panel)
	return revel.MainRouter.Reverse("Ajax.AddPanel", args).Url
}

func (_ tAjax) GetMessage(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.GetMessage", args).Url
}

func (_ tAjax) ScreenLock(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Ajax.ScreenLock", args).Url
}

func (_ tAjax) ScreenUnlock(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Ajax.ScreenUnlock", args).Url
}


type tCaptcha struct {}
var Captcha tCaptcha


func (_ tCaptcha) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.Index", args).Url
}

func (_ tCaptcha) GetCaptchaId(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Captcha.GetCaptchaId", args).Url
}


type tPublic struct {}
var Public tPublic


func (_ tPublic) Message(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Public.Message", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) Index(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Index", args).Url
}

func (_ tAdmin) Add(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Add", args).Url
}

func (_ tAdmin) Edit(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Edit", args).Url
}

func (_ tAdmin) Delete(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("Admin.Delete", args).Url
}


type tGard struct {}
var Gard tGard


func (_ tGard) Index(
		gard interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "gard", gard)
	return revel.MainRouter.Reverse("Gard.Index", args).Url
}

func (_ tGard) Add(
		gard interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "gard", gard)
	return revel.MainRouter.Reverse("Gard.Add", args).Url
}

func (_ tGard) Edit(
		gard interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "gard", gard)
	return revel.MainRouter.Reverse("Gard.Edit", args).Url
}

func (_ tGard) Delete(
		gard interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "gard", gard)
	return revel.MainRouter.Reverse("Gard.Delete", args).Url
}


type tLogs struct {}
var Logs tLogs


func (_ tLogs) Index(
		logs interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "logs", logs)
	return revel.MainRouter.Reverse("Logs.Index", args).Url
}

func (_ tLogs) DelAll(
		logs interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "logs", logs)
	return revel.MainRouter.Reverse("Logs.DelAll", args).Url
}


type tMenu struct {}
var Menu tMenu


func (_ tMenu) Index(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Index", args).Url
}

func (_ tMenu) Add(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Add", args).Url
}

func (_ tMenu) Edit(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Edit", args).Url
}

func (_ tMenu) Delete(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("Menu.Delete", args).Url
}


type tRole struct {}
var Role tRole


func (_ tRole) Index(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Index", args).Url
}

func (_ tRole) Member(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Member", args).Url
}

func (_ tRole) Add(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Add", args).Url
}

func (_ tRole) Edit(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Edit", args).Url
}

func (_ tRole) SetStatus(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.SetStatus", args).Url
}

func (_ tRole) Delete(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("Role.Delete", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Login(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Login", args).Url
}

func (_ tUser) Logout(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.Logout", args).Url
}

func (_ tUser) Left(
		menu interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "menu", menu)
	return revel.MainRouter.Reverse("User.Left", args).Url
}

func (_ tUser) EditInfo(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditInfo", args).Url
}

func (_ tUser) EditPwd(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.EditPwd", args).Url
}

func (_ tUser) AdminPanel(
		admin interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("User.AdminPanel", args).Url
}


type tPprof struct {}
var Pprof tPprof


func (_ tPprof) Profile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Profile", args).Url
}

func (_ tPprof) Symbol(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Symbol", args).Url
}

func (_ tPprof) Cmdline(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Cmdline", args).Url
}

func (_ tPprof) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Pprof.Index", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tJobs struct {}
var Jobs tJobs


func (_ tJobs) Status(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Jobs.Status", args).Url
}


