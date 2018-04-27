package report

import (
	// "fmt"
	// "log"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func List(uid string, i, j int) (map[int]map[string]string, string) {
	db := mysqldb.GetConnect()

	db.SetTable("report").Fileds("report.*").Where("report.del = 0").OrderBy("report.create_time DESC").Limit(i, j)
	ret := db.FindAll()

	total := db.SetTable("report").Fileds("COUNT(1) AS total").Where("report.del = 0").FindOne()
	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret, total[1]["total"]
}

func Del(id, uid string) int {
	db := mysqldb.GetConnect()
	i, _ := db.SetTable("report").Where("id = '" + id + "'").Update(map[string]interface{}{"del": 1})

	return i
}

func Add(param map[string]interface{}) int {
	db := mysqldb.GetConnect()

	i, _ := db.SetTable("report").Insert(param)
	return i
}

func Update(uid, id string, param map[string]interface{}) int {
	db := mysqldb.GetConnect()
	param["update_at"] = uid
	i, _ := db.SetTable("report").Where("id = '" + id + "' AND del = 0").Update(param)
	return i
}

func Info(uid, id string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("report").Where("id = '" + id + "'").FindOne()

	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret
}
