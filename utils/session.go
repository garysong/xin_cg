package utils

import "github.com/revel/revel"

func GetSession(key string, session revel.Session) string {
	if value, ok := session[key]; ok {
		return value
	}

	return ""
}
