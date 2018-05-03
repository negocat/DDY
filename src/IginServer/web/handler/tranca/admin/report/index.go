package report

import (
	// "log"
	"net/http"
	// "IginServer/lib/md5"
	// "IginServer/lib/Imartini"
	// "IginServer/lib/session"
	// "IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/lib/mygin"
	// "IginServer/lib/upload"
	"github.com/tealeg/xlsx"
	"IginServer/web/model/tranca/admin/report"
	"strconv"
	"strings"
	"time"
	"fmt"
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

	for _, i := range reports {
		i["image"] = strings.Split(i["image"], ",")[0]
	}

	c.HTML("tranca/admin/report/list.html", map[string]interface{}{
		"request": c.Request,
		"reports":   reports,
		"total":   total,
		"size":    size,
	})
}

func add(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	param := map[string]interface{}{
		"update_at":         info["id"],
		"address":        c.Request.FormValue("address"),
		"name":        c.Request.FormValue("name"),
		"phone":        c.Request.FormValue("phone"),
		"body":        c.Request.FormValue("body"),
		"image":      c.Request.FormValue("image"),
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

	p := make(map[string]interface{})

	p["image"] = c.Request.FormValue("image")
	p["body"] = c.Request.FormValue("body")
	p["address"] = c.Request.FormValue("address")
	p["name"] = c.Request.FormValue("name")
	p["phone"] = c.Request.FormValue("phone")
	report.Update(info["id"], c.Param("id"), p)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}

func down(c *mygin.IContext){
	var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    // var err error

    file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	
    row = sheet.AddRow()
    row.AddCell().Value = "ID"
    row.AddCell().Value = "姓名"
    row.AddCell().Value = "电话"
    row.AddCell().Value = "地址"
    row.AddCell().Value = "描述"
	row.AddCell().Value = "创建时间"
    row.AddCell().Value = "图片"
	
	page := 1
	size := 50

	for true {
		reports, total := report.List("0", (page-1)*size, size)
		
		for _, i := range reports {
			images := strings.Split(i["image"], ",")
			row = sheet.AddRow()
			row.AddCell().Value = i["id"]
			row.AddCell().Value = i["name"]
			row.AddCell().Value = i["phone"]
			row.AddCell().Value = i["address"]
			row.AddCell().Value = i["body"]
			row.AddCell().Value = i["create_time"]
			
			cell = row.AddCell()
			cell.SetFormula("=HYPERLINK(\"" + images[0] + "\", \"图片"+ fmt.Sprintf("%v", 1) +"\")")
			style := cell.GetStyle() 
			style.Font.Underline = true //加下划线 
			style.Font.Color = "FF0000FF"//设置字体颜色为蓝色 
			cell.SetStyle(style)

			for x, j := range images {
				if x > 0 {
					row = sheet.AddRow()
					row.AddCell()
					row.AddCell()
					row.AddCell()
					row.AddCell()
					row.AddCell()
					row.AddCell()
					cell = row.AddCell()
					cell.SetFormula("=HYPERLINK(\"" + j + "\", \"图片"+ fmt.Sprintf("%v", x+1) +"\")")
					style := cell.GetStyle() 
					style.Font.Underline = true //加下划线 
					style.Font.Color = "FF0000FF"//设置字体颜色为蓝色 
					cell.SetStyle(style)
				}
			}
			sheet.AddRow()
		}
		itotal, _ := strconv.Atoi(total)
		if  itotal < 50 {
			break
		}
		page ++
	}


	fileName := strconv.Itoa(int(time.Now().Unix()))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("content-disposition", "attachment; filename=\""+fileName+".xlsx\"")
	file.Write(c.Writer)
}