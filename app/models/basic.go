package models

//会员
//import "fmt"
import "time"

//import "xin_cg/utils"
//import "html/template"
import "github.com/revel/revel"

type Basic struct {
	Id                  int64  `xorm:"pk autoincr"`
	Code                string `xorm:"varchar(32)"`
	Lawpresent          string `xorm:"varchar(32)"`
	Lawpresentphone     string `xorm:"varchar(32)"`
	Director            string `xorm:"varchar(32)"`
	Directorphone       string `xorm:"varchar(32)"`
	Contactor           string `xorm:"varchar(32)"`
	Contactorphone      string `xorm:"varchar(32)"`
	Capital             string `xorm:"varchar(32)"`
	Fax                 string `xorm:"varchar(32)"`
	Pastalcode          string `xorm:"varchar(32)"`
	Floorarea           string `xorm:"varchar(32)"`
	Builtarea           string `xorm:"varchar(32)"`
	Outerarea           string `xorm:"varchar(32)"`
	Staffnum            string `xorm:"varchar(32)"`
	Childrennum         string `xorm:"varchar(32)"`
	Classesnum          string `xorm:"varchar(32)"`
	Securitynum         string `xorm:"varchar(32)"`
	Educationattri      string `xorm:"varchar(32)"`
	Email               string `xorm:"varchar(64)"`
	Securityinstitution string `xorm:"text"`
	Createtime          string `xorm:"DateTime"`
}

func (a *Basic) Validate(v *revel.Validation) {
	v.Required(a.Code).Message("请选择幼儿园！")
}

//根据Id获取幼儿园信息
func (a *Basic) GetById(Id int64) *Basic {
	basic := new(Basic)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false
	has, err := DB_Read.Id(Id).Get(basic)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return basic
}

func (a *Basic) GetByCode(code string) *Basic {
	basic := new(Basic)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false
	has, err := DB_Read.Where("Code=?", code).Get(basic)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return nil
	}

	return basic
}

func (a *Basic) HasBasicInfo() bool {
	basic := new(Basic)
	has, err := DB_Read.Where("code=?", a.Code).Get(basic)
	if err != nil {
		revel.WARN.Printf("错误: %v", has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	if basic.Id > 0 {
		return true
	}
	return false
}

//添加幼儿园
func (a *Basic) Save() bool {
	basic := new(Basic)

	basic.Code = a.Code
	basic.Lawpresent = a.Lawpresent
	basic.Lawpresentphone = a.Lawpresentphone
	basic.Director = a.Director
	basic.Directorphone = a.Directorphone
	basic.Contactor = a.Contactor
	basic.Contactorphone = a.Contactorphone
	basic.Capital = a.Capital
	basic.Fax = a.Fax
	basic.Pastalcode = a.Pastalcode
	basic.Floorarea = a.Floorarea
	basic.Builtarea = a.Builtarea
	basic.Outerarea = a.Outerarea
	basic.Staffnum = a.Staffnum
	basic.Childrennum = a.Childrennum
	basic.Classesnum = a.Classesnum
	basic.Securitynum = a.Securitynum
	basic.Educationattri = a.Educationattri
	basic.Email = a.Email
	basic.Securityinstitution = a.Securityinstitution
	basic.Createtime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Insert(basic)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

func (a *Basic) Edit(Id int64) bool {
	basic := new(Basic)

	basic.Lawpresent = a.Lawpresent
	basic.Lawpresentphone = a.Lawpresentphone
	basic.Director = a.Director
	basic.Directorphone = a.Directorphone
	basic.Contactor = a.Contactor
	basic.Contactorphone = a.Contactorphone
	basic.Capital = a.Capital
	basic.Fax = a.Fax
	basic.Pastalcode = a.Pastalcode
	basic.Floorarea = a.Floorarea
	basic.Builtarea = a.Builtarea
	basic.Outerarea = a.Outerarea
	basic.Staffnum = a.Staffnum
	basic.Childrennum = a.Childrennum
	basic.Classesnum = a.Classesnum
	basic.Securitynum = a.Securitynum
	basic.Educationattri = a.Educationattri
	basic.Email = a.Email
	basic.Securityinstitution = a.Securityinstitution

	has, err := DB_Write.Id(Id).Cols("Lawpresent", "Lawpresentphone", "Director", "Directorphone", "Contactor", "Contactorphone", "Capital", "Fax", "Pastalcode", "Floorarea", "Builtarea", "Outerarea", "Staffnum", "Childrennum", "Classesnum", "Securitynum", "EducationAttri", "Email", "Securityinstitution").Update(basic)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}
