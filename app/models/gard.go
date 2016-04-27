package models

//管理员表
import "time"
import "xin_cg/utils"
import "html/template"
import "github.com/revel/revel"

type Gard struct {
	Id           int64  `xorm:"pk autoincr"`
	Dict         string `xorm:"varchar(32)"`
	Street       string `xorm:"varchar(32)"`
	Childrengard string `xorm:"varchar(64)"`
	Code         string `xorm:"varchar(10)"`
	Createtime   string `xorm:"DateTime"`
}

func (a *Gard) Validate(v *revel.Validation) {
	v.Required(a.Dict).Message("请输入区名称！")
	//v.Required(a.Street).Message("请输入街道名称！")
	//v.Required(a.Childrengard).Message("请输入幼儿园名称！")
	v.Required(a.Code).Message("请输入编码！")
}

//获取幼儿园列表
func (a *Gard) GetByAll(Page int64, Perpage int64) ([]*Gard, template.HTML) {
	gard_list := []*Gard{}

	// 查询总数
	gard := new(Gard)
	Total, err := DB_Read.Count(gard)

	if err != nil {
		revel.WARN.Printf("错误:%v", err)
	}

	// 分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Gard/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//
	DB_Read.Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Find(&gard_list)

	return gard_list, pages

}

//根据Id获取幼儿园信息
func (a *Gard) GetById(Id int64) *Gard {
	gard := new(Gard)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false
	has, err := DB_Read.Id(Id).Get(gard)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return gard
}

//查找幼儿园编码是否存在
func (a *Gard) HasGard() bool {
	gard := new(Gard)
	has, err := DB_Read.Where("code=?", a.Code).Get(gard)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if gard.Id > 0 {
		return true
	}
	return false
}

//添加幼儿园
func (a *Gard) Save() bool {
	gard := new(Gard)
	gard.Dict = a.Dict
	gard.Street = a.Street
	gard.Childrengard = a.Childrengard
	gard.Code = a.Code

	gard.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(gard)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//修改幼儿园信息
func (a *Gard) Edit(Id int64) bool {
	gard := new(Gard)

	if len(a.Dict) > 0 {
		gard.Dict = a.Dict
	}

	if len(a.Street) > 0 {
		gard.Street = a.Street
	}

	if len(a.Childrengard) > 0 {
		gard.Childrengard = a.Childrengard
	}

	if len(a.Code) > 0 {
		gard.Code = a.Code
	}

	has, err := DB_Write.Id(Id).Cols("Dict", "Street", "Childrengard", "Code").Update(gard)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

// 删除幼儿园
func (a *Gard) DelByID(Id int64) bool {
	gard := new(Gard)

	has, err := DB_Write.Id(Id).Delete(gard)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//获取区名称
func (a *Gard) GetDictList() [5]string {
	dict_list := [5]string{"罗湖", "南山", "福田", "宝安", "龙岗"}
	return dict_list
}
