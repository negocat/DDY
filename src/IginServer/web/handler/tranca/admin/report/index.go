package report

import (
	// "log"
	"net/http"
	// "IginServer/lib/md5"
	// "IginServer/lib/Imartini"
	// "IginServer/lib/session"
	"IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/lib/mygin"
	"IginServer/lib/upload"
	"IginServer/web/model/tranca/admin/report"
	"strconv"
	"time"
)

func list(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	var page, size int
	page, _ = strconv.Atoi(c.Request.FormValue("page"))
	if page <= 0 {
		page = 1
	}
	size = 10
	reports, total := report.List(info["id"], (page-1)*size, size)

	c.HTML("tranca/admin/report/list.html", map[string]interface{}{
		"request": c.Request,
		"reports":   reports,
		"total":   total,
		"size":    size,
	})
}

func add(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "image", 0, 0, info["id"])

	param := map[string]interface{}{
		"update_at":         info["id"],
		"address":        c.Request.FormValue("address"),
		"name":        c.Request.FormValue("name"),
		"phone":        c.Request.FormValue("phone"),
		"body":        c.Request.FormValue("body"),
		"image":      conf.GetString("config", "IMAGE_HOST") + path,
		"update_time": int(time.Now().Unix()),
		"create_time": int(time.Now().Unix()),
	}
	report.Add(param)
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))

}

func del(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	//0：删除失败，1：删除成功
	// return ddy.DelWord(c.Param("id"), info["id"])
	if report.Del(c.Param("id"), info["id"]) == 0 {
		c.String(200, "0")
	} else {
		c.String(200, "1")
	}

}

func info(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	reports := report.Info(info["id"], c.Param("id"))

	if len(reports) > 0 {
		R.Write(c.Writer, map[string]interface{}{"report": reports[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}

func update(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "image", 0, 0, info["id"])
	p := make(map[string]interface{})

	if path != "" {
		p["image"] = conf.GetString("config", "IMAGE_HOST") + path
	}
	p["body"] = c.Request.FormValue("body")
	p["address"] = c.Request.FormValue("address")
	p["name"] = c.Request.FormValue("name")
	p["phone"] = c.Request.FormValue("phone")
	report.Update(info["id"], c.Param("id"), p)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}
