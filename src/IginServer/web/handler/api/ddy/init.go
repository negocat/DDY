package ddy

import (
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/web/handler/api/auth"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("api/cartoon")
	r.GET("list", mygin.Handler(cartoon))
	r.GET("info", mygin.Handler(cartoonInfo))
	r.GET("episodesInfo", mygin.Handler(cartoonEpisodesInfo))

	r.GET("class", mygin.Handler(classList))
}
