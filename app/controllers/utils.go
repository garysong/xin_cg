package controllers

import "strconv"
import "github.com/revel/revel"
import "xin_cg/app/models"

func GetCode(c interface{}) string {

	r := c.(*revel.Controller)

	if UserID, ok := r.Session["UserID"]; ok {
		UserID, err := strconv.ParseInt(UserID, 10, 64)
		if err != nil {
			revel.WARN.Println(err)
		}

		admin := new(models.Admin)

		admin_info := admin.GetById(UserID)

		return admin_info.Code

	}

	return ""
}
