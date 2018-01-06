package ddy

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
	"IginServer/web/model/tranca/admin/ddy"
	"strconv"
	"time"
)

func banner(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	class := ddy.Class(info["id"])

	var page, size int
	page, _ = strconv.Atoi(c.Request.FormValue("page"))
	if page <= 0 {
		page = 1
	}
	size = 10
	words, total := ddy.Word(info["id"], (page-1)*size, size)

	c.HTML("tranca/admin/ddy/banner.html", map[string]interface{}{
		"request": c.Request,
		"class":   class,
		"words":   words,
		"total":   total,
		"size":    size,
	})
}

func addbanner(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])

	param := map[string]interface{}{
		"uid":         info["id"],
		"name":        c.Request.FormValue("wordname"),
		"classid":     c.Request.FormValue("classid"),
		"imgurl":      conf.GetString("config", "IMAGE_HOST") + path,
		"remarks":     c.Request.FormValue("remarks"),
		"wordbody":    c.Request.FormValue("wordbody"),
		"create_time": int(time.Now().Unix()),
	}
	ddy.AddWord(param)
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))

}

func delbanner(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	//0：删除失败，1：删除成功
	// return ddy.DelWord(c.Param("id"), info["id"])
	if ddy.DelWord(c.Param("id"), info["id"]) == 0 {
		c.String(200, "0")
	} else {
		c.String(200, "1")
	}

}

func bannerinfo(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	words := ddy.Wordinfo(info["id"], c.Param("id"))

	if len(words) > 0 {
		R.Write(c.Writer, map[string]interface{}{"words": words[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}

func updatebanner(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])
	p := make(map[string]interface{})

	if path != "" {
		p["imgurl"] = conf.GetString("config", "IMAGE_HOST") + path
	}
	p["classid"] = c.Request.FormValue("classid")
	p["name"] = c.Request.FormValue("wordname")
	p["remarks"] = c.Request.FormValue("remark")
	p["wordbody"] = c.Request.FormValue("wordbody")
	ddy.UpdateWord(info["id"], c.Param("id"), p)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}
