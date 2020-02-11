package models

import (
	"ginexample/common/mysql"
	"ginexample/utils/log"
)

func GetMysqlVersion() *string {
	version := new(string)
	err := mysql.SQLX_DB.Get(version, "select version()")
	if err != nil {
		log.Warnf("err: %s", err)
		return version
	}
	return version
}
