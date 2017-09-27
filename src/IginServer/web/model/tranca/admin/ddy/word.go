package ddy

import (
	// "fmt"
	// "log"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func Word(uid string, i, j int) (map[int]map[string]string, string) {
	db := mysqldb.GetConnect()

	db.SetTable("i_goods_word").Fileds("i_goods_word.*", "i_goods_class.name AS classname").Where("i_goods_word.del = 0")
	db.LeftJoin("i_goods_class", "i_goods_word.classid = i_goods_class.id").OrderBy("i_goods_word.create_time DESC").Limit(i, j)
	ret := db.FindAll()

	total := db.SetTable("i_goods_word").Fileds("COUNT(1) AS total").Where("i_goods_word.del = 0").FindOne()
	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret, total[1]["total"]
}

func DelWord(id, uid string) int {
	db := mysqldb.GetConnect()
	i, _ := db.SetTable("i_goods_word").Where("id = '" + id + "'").Update(map[string]interface{}{"del": 1})

	return i
}

func AddWord(param map[string]interface{}) int {
	db := mysqldb.GetConnect()

	i, _ := db.SetTable("i_goods_word").Insert(param)
	return i
}

func UpdateWord(uid, id string, param map[string]interface{}) int {
	db := mysqldb.GetConnect()
	param["uid"] = uid
	i, _ := db.SetTable("i_goods_word").Where("id = '" + id + "' AND del = 0").Update(param)
	return i
}

func Wordinfo(uid, id string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("i_goods_word").Where("id = '" + id + "'").FindOne()

	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret
}
