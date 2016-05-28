package models

import "time"

import "xin_cg/utils"
import "html/template"
import "github.com/revel/revel"

type Equipment struct {
	Id         int64  `xorm:"pk autoincr"`
	Code       string `xorm:"varchar(32)"`
	Name       string `xorm:"varchar(32)"`
	Model      string `xorm:"varchar(64)"`
	Num        int64  `xorm:"int(11)"`
	Status     string `xorm:"varchar(64)"`
	Createtime string `xorm:"DateTime"`
}

func (a *Equipment) Validate(v *revel.Validation) {
	v.Required(a.Code).Message("请选择幼儿园！")
}

func (a *Equipment) GetByCode(code string, where map[string]string, Page int64, Perpage int64) ([]*Equipment, template.HTML) {
	equipment_list := []*Equipment{}

	if len(code) > 0 {
		equipment := new(Equipment)
		Total, err := DB_Read.Where("code=?", code).Count(equipment)
		if err != nil {
			revel.WARN.Printf("错误：%v", err)
		}

		Pager := new(utils.Page)
		Pager.SubPage_link = "/Basic/"
		Pager.Nums = Total
		Pager.Perpage = Perpage
		Pager.Current_page = Page
		Pager.SubPage_type = 2
		pages := Pager.Show()

		DB_Read.Where("code=?", code).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&equipment_list)

		return equipment_list, pages
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Basic/"
	Pager.Nums = 0
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	return equipment_list, pages
}

//根据Id获取消防设备信息
func (a *Equipment) GetById(Id int64) *Equipment {
	equipment := new(Equipment)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(equipment)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return equipment
}

//添加消防设备
func (a *Equipment) Save() bool {

	equipment := new(Equipment)
	equipment.Code = a.Code
	equipment.Name = a.Name
	equipment.Model = a.Model
	equipment.Num = a.Num
	equipment.Status = a.Status
	equipment.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(equipment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑消防设备
func (a *Equipment) Edit(Id int64) bool {

	equipment := new(Equipment)

	if len(a.Name) > 0 {
		equipment.Name = a.Name
	}

	if len(a.Model) > 0 {
		equipment.Model = a.Model
	}

	if a.Num > 0 {
		equipment.Num = a.Num
	}

	if len(a.Status) > 0 {
		equipment.Status = a.Status
	}

	has, err := DB_Write.Id(Id).Cols("name", "model", "num", "status").Update(equipment)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//删除消防设备
func (a *Equipment) DelByID(Id int64) bool {

	equipment := new(Equipment)

	has, err := DB_Write.Id(Id).Delete(equipment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Println("错误: %v", err)
		return false
	}

	return true
}
