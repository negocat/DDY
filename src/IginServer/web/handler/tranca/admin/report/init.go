package report

import (
	// "IginServer/framework/ueditor"
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/web/handler/tranca/admin/auth"
	// "IginServer/web"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("admin", mygin.Handler(auth.Auth))
	r.GET("/report/list", mygin.Handler(list))
	r.POST("/report/list", mygin.Handler(add))
	r.GET("/report/del/:id", mygin.Handler(del))
	r.GET("/report/info/:id", mygin.Handler(info))
	r.POST("/report/update/:id", mygin.Handler(update))
}
