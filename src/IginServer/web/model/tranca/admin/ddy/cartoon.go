package ddy

import (
	"fmt"
	// "log"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func Cartoon(uid string, i, j int) (map[int]map[string]string, string) {
	db := mysqldb.GetConnect()

	db.SetTable("cartoon").Fileds("cartoon.*", "i_goods_class.name AS classname").Where("cartoon.del = 0")
	db.LeftJoin("i_goods_class", "cartoon.classid = i_goods_class.id").OrderBy("cartoon.create_time DESC").Limit(i, j)
	ret := db.FindAll()

	total := db.SetTable("cartoon").Fileds("COUNT(1) AS total").Where("cartoon.del = 0").FindOne()
	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret, total[1]["total"]
}

func DelCartoon(id, uid string) int {
	db := mysqldb.GetConnect()
	i, _ := db.SetTable("cartoon").Where("id = '" + id + "'").Update(map[string]interface{}{"del": 1})

	return i
}

func AddCartoon(param map[string]interface{}) int {
	db := mysqldb.GetConnect()

	i, err := db.SetTable("cartoon").Insert(param)
	fmt.Println(err)
	return i
}

func UpdateCartoon(uid, id string, param map[string]interface{}) int {
	db := mysqldb.GetConnect()
	param["uid"] = uid
	i, _ := db.SetTable("cartoon").Where("id = '" + id + "' AND del = 0").Update(param)
	return i
}

func CartoonInfo(uid, id string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("cartoon").Where("id = '" + id + "'").FindOne()

	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret
}
