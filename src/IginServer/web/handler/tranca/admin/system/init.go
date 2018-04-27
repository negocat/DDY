package system

import (
	// "IginServer/framework/ueditor"
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/web/handler/tranca/admin/auth"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("admin", mygin.Handler(auth.Auth))
	r.GET("system/password", mygin.Handler(password))
	r.POST("system/password", mygin.Handler(setPassword))
	r.GET("system/logout", mygin.Handler(logout))
}
