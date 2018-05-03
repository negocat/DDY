package report

import (
	// "IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/conf"
	"IginServer/framework/API/R"
	// "IginServer/web/handler/api/sms"
	"IginServer/web/model/tranca/admin/report"
	// "encoding/json"
	// w "IginServer/framework/weixin"
	// "IginServer/lib/md5"
	// "IginServer/lib/other"
	// "IginServer/lib/redis"
	// "IginServer/lib/session"
	"IginServer/lib/upload"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "net/http"
	"time"
)


func Add(c *gin.Context) {

	path := ""
	imgs := c.Request.Form["image"]
	for _, img := range imgs {
		imgPath, _ := upload.ImgStrBase64Save(img, 0, 0, "image")
		path += conf.GetString("config", "IMAGE_HOST") + imgPath + ","
	}

	param := map[string]interface{}{
		"update_at":         0,
		"address":        c.Request.FormValue("address"),
		"name":        c.Request.FormValue("name"),
		"phone":        c.Request.FormValue("phone"),
		"body":        c.Request.FormValue("body"),
		"image":      path,
		"update_time": int(time.Now().Unix()),
		"create_time": int(time.Now().Unix()),
	}
	report.Add(param)
	R.Success(c.Writer, "")
}
