package system

import (
	// "fmt"
	// "log"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func SetPassword(password, uid string) {
	db := mysqldb.GetConnect()

	param := map[string]interface{}{"passwd": password}

	db.SetTable("admin_auth").Where("id = " + uid).Update(param)
}