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

func cartoon(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	class := ddy.Class(info["id"])

	var page, size int
	page, _ = strconv.Atoi(c.Request.FormValue("page"))
	if page <= 0 {
		page = 1
	}
	size = 10
	cartoons, total := ddy.Cartoon(info["id"], (page-1)*size, size, nil)

	c.HTML("tranca/admin/ddy/cartoon.html", map[string]interface{}{
		"request": c.Request,
		"class":   class,
		"cartoons":   cartoons,
		"total":   total,
		"size":    size,
	})
}

func addcartoon(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])

	param := map[string]interface{}{
		"uid":         info["id"],
		"name":        c.Request.FormValue("name"),
		"classid":     c.Request.FormValue("classid"),
		"author":      c.Request.FormValue("author"),
		"imgurl":      conf.GetString("config", "IMAGE_HOST") + path,
		"remarks":     c.Request.FormValue("remarks"),
		"abstract":    c.Request.FormValue("abstract"),
		"create_time": int(time.Now().Unix()),
		"update_time": int(time.Now().Unix()),
	}
	ddy.AddCartoon(param)
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))

}

func delcartoon(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	//0：删除失败，1：删除成功
	// return ddy.DelWord(c.Param("id"), info["id"])
	if ddy.DelCartoon(c.Param("id"), info["id"]) == 0 {
		c.String(200, "0")
	} else {
		c.String(200, "1")
	}

}

func cartooninfo(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	cartoons := ddy.CartoonInfo(info["id"], c.Param("id"))

	if len(cartoons) > 0 {
		R.Write(c.Writer, map[string]interface{}{"cartoons": cartoons[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}

func updatecartoon(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])
	p := make(map[string]interface{})

	if path != "" {
		p["imgurl"] = conf.GetString("config", "IMAGE_HOST") + path
	}
	p["uid"] = info["id"]
	p["name"] = c.Request.FormValue("name")
	p["classid"] = c.Request.FormValue("classid")
	p["author"] = c.Request.FormValue("author")
	p["remarks"] = c.Request.FormValue("remarks")
	p["abstract"] = c.Request.FormValue("abstract")
	p["update_time"] = int(time.Now().Unix())
	ddy.UpdateCartoon(info["id"], c.Param("id"), p)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}

func cartoonEpisodes(c *mygin.IContext) {
	// info := c.GetSession("admin_info").(map[string]string)

	cartoonEpisodes, total := ddy.CartoonEpisodes(c.Param("id"))

	c.HTML("tranca/admin/ddy/cartoonEpisodes.html", map[string]interface{}{
		"request": c.Request,
		"cartoonEpisodes":   cartoonEpisodes,
		"total":   total,
	})
}

func addCartoonEpisodes(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	// path, _ := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])

	param := map[string]interface{}{
		"uid":         info["id"],
		"title":        c.Request.FormValue("title"),
		"sort":     c.Request.FormValue("sort"),
		"imglist":  c.Request.FormValue("imgs"),
		"cartoon_id":      c.Param("id"),
		"createtime": int(time.Now().Unix()),
	}
	ddy.AddCartoonEpisodes(param)
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))

}

func delCartoonEpisodes(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	//0：删除失败，1：删除成功
	// return ddy.DelWord(c.Param("id"), info["id"])
	if ddy.DelCartoonEpisodes(c.Param("id"), info["id"]) == 0 {
		c.String(200, "0")
	} else {
		c.String(200, "1")
	}

}

func cartoonEpisodesInfo(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	cartoons := ddy.CartoonEpisodesInfo(info["id"], c.Param("id"))

	if len(cartoons) > 0 {
		R.Write(c.Writer, map[string]interface{}{"cartoonEpisodes": cartoons[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}

func updateCartoonEpisodes(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	param := map[string]interface{}{
		"uid":         info["id"],
		"title":        c.Request.FormValue("title"),
		"sort":     c.Request.FormValue("sort"),
		"imglist":  c.Request.FormValue("imgs"),
		"createtime": int(time.Now().Unix()),
	}
	ddy.UpdateCartoonEpisodes(info["id"], c.Param("id"), param)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}