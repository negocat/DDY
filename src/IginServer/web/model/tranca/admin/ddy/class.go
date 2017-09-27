package ddy

import (
	"fmt"
	// "log"
	"IginServer/lib/mysqldb"
	"strconv"
)

func Class(uid string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("i_goods_class").OrderBy("topid, list").FindAll()

	// top := []map[string]string{}
	// for _, i := range ret {
	// 	if i["level"] == "1" {
	// 		top = append(top, i)
	// 	}
	// }
	// fmt.Println(top)

	return ret
}

func DelClass(id, uid string) string {
	db := mysqldb.GetConnect()

	// info := db.SetTable("i_goods_class").Where("uid = '" + uid + "' and id = '" + id + "'").FindOne()
	upper := db.SetTable("i_goods_class").Where("upperid = '" + id + "'").FindOne()

	if len(upper) > 0 {
		return "-1"
	} else {
		goods := db.SetTable("i_goods").Where(" class = '" + id + "'").FindOne()
		if len(goods) > 0 {
			return "0"
		}
	}
	db.SetTable("i_goods_class").Delete(" id = '" + id + "'")

	return "1"
}

func AddClass(param map[string]interface{}) {
	db := mysqldb.GetConnect()
	param["list"] = ""
	var sql string

	info := db.SetTable("i_goods_class").Where("id = '" + param["upperid"].(string) + "'").FindOne()
	if len(info) > 0 {
		stmp, _ := strconv.Atoi(info[1]["level"])
		stmp += 1
		param["level"] = strconv.Itoa(stmp)
		param["topid"] = info[1]["topid"]
		param["list"] = info[1]["list"]
		x, _ := db.SetTable("i_goods_class").Insert(param)

		sql = "UPDATE i_goods_class SET list = CONCAT(list, id, ',') where id = '%v'"
		sql = fmt.Sprintf(sql, x)
	} else {
		param["level"] = 1
		delete(param, "upperid")
		x, _ := db.SetTable("i_goods_class").Insert(param)

		sql = "UPDATE i_goods_class SET topid = id, list = CONCAT(list, id, ',') where id = '%v'"
		sql = fmt.Sprintf(sql, x)
	}
	db.Query(sql)
}

func UpdateClass(uid, id string, param map[string]interface{}) (int, error) {
	db := mysqldb.GetConnect()
	param["uid"] = uid
	return db.SetTable("i_goods_class").Where("id = '" + id + "'").Update(param)
}

func Classinfo(uid, id string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("i_goods_class").Where(" id = '" + id + "'").FindOne()

	// top := []map[string]string{}
	// for _, i := range ret {
	// 	if i["level"] == "1" {
	// 		top = append(top, i)
	// 	}
	// }
	// fmt.Println(top)

	return ret
}
