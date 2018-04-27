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
	// "IginServer/lib/md5"
	// "IginServer/lib/other"
	// "IginServer/lib/upload"
	// "IginServer/web/model/tranca/admin/system"
	// "IginServer/web/model/tranca/product"
	// "strconv"
	// "strings"
	// "time"
)

func logout(c *mygin.IContext) {
	c.DelSession("admin_info")

	c.Redirect(http.StatusFound, "/admin/")
}