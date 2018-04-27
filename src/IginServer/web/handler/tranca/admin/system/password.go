package system

import (
	"IginServer/lib/mygin"
	// "log"
	// "github.com/go-martini/martini"
	// "fmt"
	// "IginServer/conf"
	"net/http"
	// "IginServer/lib/Imartini"
	// "IginServer/framework/API/R"
	// "IginServer/lib/cookie"
	"IginServer/lib/md5"
	// "IginServer/lib/other"
	// "IginServer/lib/upload"
	"IginServer/web/model/tranca/admin/system"
	// "IginServer/web/model/tranca/product"
	// "strconv"
	// "strings"
	// "time"
)

func password(c *mygin.IContext) {

	c.HTML("tranca/admin/system/password.html", map[string]interface{}{})
}

func setPassword(c *mygin.IContext) {
	password := c.Request.FormValue("password")
	if len(password) <= 0 {
		c.Redirect(http.StatusFound, c.Request.Header.Get("Referer") + "?msg=密码不能为空")
		return
	}

	info := c.GetSession("admin_info").(map[string]string)
	password = md5.Md5(password)
	system.SetPassword(password, info["id"])
	
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer") + "?msg=密码修改成功")
}