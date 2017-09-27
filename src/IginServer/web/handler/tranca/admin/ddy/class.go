package ddy

import (
	"log"
	"net/http"
	// "IginServer/lib/md5"
	// "IginServer/lib/Imartini"
	// "IginServer/lib/session"
	"IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/lib/mygin"
	"IginServer/lib/upload"
	"IginServer/web/model/tranca/admin/ddy"
	"time"
)

var (
	wlogo   = conf.GetInt("app", "logowidth")
	hlogo   = conf.GetInt("app", "logoheight")
	wbanner = conf.GetInt("app", "bannerwidth")
	hbanner = conf.GetInt("app", "bannerheight")
)

func Class(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	class := ddy.Class(info["id"])

	c.HTML("tranca/admin/ddy/class.html", map[string]interface{}{
		"class": class,
	})
}

func delclass(c *mygin.IContext) {
	// sess := session.GetAll(req)
	info := c.GetSession("admin_info").(map[string]string)

	id := c.Param("id")

	//-1：不是末级科目，0：类别下还有商品，1：删除成功
	c.String(200, ddy.DelClass(id, info["id"]))
}

func addclass(c *mygin.IContext) {
	// sess := session.GetAll(req)
	info := c.GetSession("admin_info").(map[string]string)

	path, _ := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])
	// if err != nil && err.Error() == "-1" {
	// 	res.Write([]byte("上传的图片类型不正确"))
	// 	return
	// }
	classname := c.Request.FormValue("classname")
	upperid := c.Request.FormValue("upperid")
	remark := c.Request.FormValue("remark")
	imgurl := conf.GetString("config", "IMAGE_HOST") + path

	param := map[string]interface{}{
		"uid":         info["id"],
		"name":        classname,
		"upperid":     upperid,
		"imgurl":      imgurl,
		"remarks":     remark,
		"create_time": int(time.Now().Unix()),
	}
	ddy.AddClass(param)
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}

func updateclass(c *mygin.IContext) {
	// sess := session.GetAll(req)
	info := c.GetSession("admin_info").(map[string]string)

	path, err := upload.ImgSave(c.Request, "imgurl", wlogo, hlogo, info["id"])
	if err != nil && err.Error() == "-1" {
		c.String(200, "上传的图片类型不正确")
	} else if err != nil {
		log.Printf("%v:%s\n", err, path)
	}
	p := make(map[string]interface{})
	if path != "" {
		p["imgurl"] = conf.GetString("config", "IMAGE_HOST") + path
	}
	p["name"] = c.Request.FormValue("classname")
	p["remarks"] = c.Request.FormValue("remark")
	_, err = ddy.UpdateClass(info["id"], c.Param("id"), p)
	if err != nil {
		log.Printf("%v\n", err)
	}
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}

func getclassinfo(c *mygin.IContext) {
	// sess := session.GetAll(req)
	info := c.GetSession("admin_info").(map[string]string)

	class := ddy.Classinfo(info["id"], c.Param("id"))

	if len(class) > 0 {
		R.Write(c.Writer, map[string]interface{}{"class": class[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}
