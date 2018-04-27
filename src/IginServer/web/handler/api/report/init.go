package report

import (
	"IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/web/handler/api/auth"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("api/report")
	r.POST("add", Add)
}
