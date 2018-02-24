package ddy

import (
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/web"
	"IginServer/web/handler/tranca/admin/auth"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("admin", mygin.Handler(auth.Auth))
	r.GET("/ddy/class", mygin.Handler(Class))
	r.GET("/ddy/delclass/:id", mygin.Handler(delclass))
	r.POST("/ddy/class", mygin.Handler(addclass))
	r.POST("/ddy/updateclass/:id", mygin.Handler(updateclass))
	r.GET("/ddy/class/:id", mygin.Handler(getclassinfo))

	r.GET("/ddy/word", mygin.Handler(word))
	r.GET("/ddy/delword/:id", mygin.Handler(delword))
	r.POST("/ddy/word", mygin.Handler(addword))
	r.POST("/ddy/updateword/:id", mygin.Handler(updateword))
	r.GET("/ddy/wordinfo/:id", mygin.Handler(wordinfo))

	r.GET("/ddy/banner", mygin.Handler(banner))

	r.GET("/ddy/cartoon", mygin.Handler(cartoon))
	r.POST("/ddy/cartoon", mygin.Handler(addcartoon))
	r.GET("/ddy/delcartoon/:id", mygin.Handler(delcartoon))
	r.POST("/ddy/updatecartoon/:id", mygin.Handler(updatecartoon))
	r.GET("/ddy/cartooninfo/:id", mygin.Handler(cartooninfo))
	
	r.GET("/ddy/cartoonEpisodes/:id", mygin.Handler(cartoonEpisodes))
	r.POST("/ddy/cartoonEpisodes/:id", mygin.Handler(addCartoonEpisodes))
	r.GET("/ddy/delCartoonEpisodes/:id", mygin.Handler(delCartoonEpisodes))
	r.POST("/ddy/updateCartoonEpisodes/:id", mygin.Handler(updateCartoonEpisodes))
	r.GET("/ddy/cartoonEpisodesInfo/:id", mygin.Handler(cartoonEpisodesInfo))
}
