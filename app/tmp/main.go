// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers8 "github.com/revel/revel/modules/jobs/app/controllers"
	_ "github.com/revel/revel/modules/jobs/app/jobs"
	controllers5 "github.com/revel/revel/modules/pprof/app/controllers"
	controllers6 "github.com/revel/revel/modules/static/app/controllers"
	_ "github.com/revel/revel/modules/testrunner/app"
	controllers7 "github.com/revel/revel/modules/testrunner/app/controllers"
	_ "xin_cg/app"
	controllers "xin_cg/app/controllers"
	controllers0 "xin_cg/app/controllers/Basic"
	controllers1 "xin_cg/app/controllers/Panel"
	controllers2 "xin_cg/app/controllers/Public"
	controllers3 "xin_cg/app/controllers/Setting"
	controllers4 "xin_cg/app/controllers/User"
	models "xin_cg/app/models"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					52: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Main",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					108: []string{ 
						"title",
						"admin_info",
						"system_info",
						"panel_list",
					},
					110: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Basic)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "basic", Type: reflect.TypeOf((**models.Basic)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					32: []string{ 
						"title",
					},
					35: []string{ 
						"title",
						"gardname",
					},
					39: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Info",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "basic", Type: reflect.TypeOf((**models.Basic)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					172: []string{ 
						"title",
						"gardname",
					},
					177: []string{ 
						"title",
						"admin_info",
						"basic_info",
						"gardname",
					},
					180: []string{ 
						"title",
						"admin_info",
						"gardname",
					},
					182: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Equipment)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "equipment", Type: reflect.TypeOf((**models.Equipment)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					34: []string{ 
						"title",
						"equipment_list",
						"pages",
						"gardname",
					},
					37: []string{ 
						"title",
						"equipment_list",
						"pages",
						"gardname",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "equipment", Type: reflect.TypeOf((**models.Equipment)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					53: []string{ 
						"title",
					},
					56: []string{ 
						"title",
						"gardname",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "equipment", Type: reflect.TypeOf((**models.Equipment)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					127: []string{ 
						"title",
						"equipment_info",
						"gardname",
					},
					130: []string{ 
						"title",
					},
					133: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "equipment", Type: reflect.TypeOf((**models.Equipment)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Panel)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					13: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Ajax)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "GetCaptcha",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Pos",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin_panel", Type: reflect.TypeOf((**models.Admin_Panel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DelPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin_panel", Type: reflect.TypeOf((**models.Admin_Panel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin_panel", Type: reflect.TypeOf((**models.Admin_Panel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetMessage",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ScreenLock",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ScreenUnlock",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Captcha)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetCaptchaId",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers2.Public)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Message",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					15: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					30: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					34: []string{ 
						"title",
						"admin_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					50: []string{ 
						"title",
						"role_list",
						"gname_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					227: []string{ 
						"title",
						"admin_info",
						"role_list",
						"gname_list",
					},
					229: []string{ 
						"title",
						"role_list",
						"gname_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Gard)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "gard", Type: reflect.TypeOf((**models.Gard)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"title",
						"gard_list",
						"pages",
					},
					33: []string{ 
						"title",
						"gard_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "gard", Type: reflect.TypeOf((**models.Gard)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
						"title",
						"gard_list",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "gard", Type: reflect.TypeOf((**models.Gard)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					132: []string{ 
						"title",
						"gard_info",
						"gard_list",
					},
					134: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "gard", Type: reflect.TypeOf((**models.Gard)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Logs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "logs", Type: reflect.TypeOf((**models.Logs)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					28: []string{ 
						"title",
						"logs_list",
						"where",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "DelAll",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "logs", Type: reflect.TypeOf((**models.Logs)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Menu)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					26: []string{ 
						"title",
						"menus",
					},
					28: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					58: []string{ 
						"title",
						"menus",
						"Id",
					},
					60: []string{ 
						"title",
						"Id",
					},
					75: []string{ 
						"title",
						"menus",
					},
					77: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					212: []string{ 
						"title",
						"menus",
						"menu_info",
					},
					214: []string{ 
						"title",
						"menu_info",
					},
					231: []string{ 
						"title",
						"menus",
					},
					233: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers3.Role)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					26: []string{ 
						"title",
						"role_list",
						"pages",
					},
					30: []string{ 
						"title",
						"role_list",
						"pages",
					},
				},
			},
			&revel.MethodType{
				Name: "Member",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					61: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					65: []string{ 
						"title",
						"admin_list",
						"pages",
					},
					69: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Add",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					94: []string{ 
						"title",
						"tree",
					},
					96: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "Edit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					200: []string{ 
						"title",
						"role_info",
						"tree",
						"Id",
					},
					202: []string{ 
						"title",
						"role_info",
						"Id",
					},
					220: []string{ 
						"title",
						"tree",
					},
					222: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "SetStatus",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**models.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers4.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					22: []string{ 
						"title",
						"CaptchaId",
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Left",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "menu", Type: reflect.TypeOf((**models.Menu)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					190: []string{ 
						"title",
						"left_menu",
					},
					192: []string{ 
						"title",
					},
					210: []string{ 
						"title",
						"left_menu",
					},
					212: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "EditInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					230: []string{ 
						"title",
						"admin_info",
					},
					232: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "EditPwd",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					328: []string{ 
						"title",
						"admin_info",
					},
					330: []string{ 
						"title",
					},
				},
			},
			&revel.MethodType{
				Name: "AdminPanel",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((**models.Admin)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					410: []string{ 
						"title",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers5.Pprof)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Profile",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Symbol",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Cmdline",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers6.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers7.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					78: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers8.Jobs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Status",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					19: []string{ 
						"entries",
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"xin_cg/app/models.(*Admin).Validate": { 
			33: "a.Username",
			34: "a.Username",
			48: "a.Email",
			49: "a.Email",
			63: "a.Password",
			64: "a.Password",
		},
		"xin_cg/app/models.(*Basic).Validate": { 
			38: "a.Code",
		},
		"xin_cg/app/models.(*Equipment).Validate": { 
			20: "a.Code",
		},
		"xin_cg/app/models.(*Gard).Validate": { 
			19: "a.Dict",
			22: "a.Code",
		},
		"xin_cg/app/models.(*Menu).Validate": { 
			22: "menu.Name",
			23: "menu.Name",
			24: "menu.Pid",
			25: "menu.Url",
			26: "menu.Order",
		},
		"xin_cg/app/models.(*Password).ValidatePassword": { 
			69: "P.Password",
			70: "P.PasswordConfirm",
			72: "P.Password",
			73: "P.Password",
		},
	}
	revel.TestSuites = []interface{}{ 
	}

	revel.Run(*port)
}
