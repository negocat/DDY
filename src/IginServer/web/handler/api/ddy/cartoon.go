package ddy

import (
	// "log"
	"IginServer/lib/mysqldb"
	// "IginServer/lib/md5"
	// "IginServer/lib/Imartini"
	// "IginServer/lib/session"
	// "IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/lib/mygin"
	"IginServer/web/model/tranca/admin/ddy"
	"strconv"
)


/**
 * @api {get} /api/cartoon/list 获取漫画列表
 * @apiName 获取漫画列表
 * @apiGroup 漫画
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {int} page 页数，从1开始
 * @apiParam {int} size 条数，默认10
 * @apiParam {int} classId 分类id
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"data":[
 *    {
 *   	 "abstract":"",   //摘要
 *   	 "author":"2",    //作者
 *   	 "classid":"7",
 *   	 "classname":"图片",
 *   	 "create_time":"1517542845",
 *   	 "del":"0",
 *   	 "id":"3",
 *   	 "imgurl":"http://121.41.116.104:8003/",         // 封面
 *   	 "name":"1",                     // 漫画名称
 *   	 "remarks":"",                   // 备注
 *   	 "uid":"4",
 *   	 "update_time":"1517542845"
 *    }],"status":200}
 *
 *
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 *
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://121.41.116.104:8003/api/cartoon/list
 */
func cartoon(c *mygin.IContext) {

	var page, size int
	page, _ = strconv.Atoi(c.Request.FormValue("page"))
	size, _ = strconv.Atoi(c.Request.FormValue("size"))
	classId := c.Request.FormValue("classId")
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	data, _ := ddy.Cartoon("0", (page-1)*size, size, map[string]interface{}{"classid": classId})

	R.Success(c.Writer, mysqldb.List(data))
}


/**
 * @api {get} /api/cartoon/info 获取漫画详情
 * @apiName 获取漫画详情
 * @apiGroup 漫画
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {int} id 漫画id
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"data":
 *    {
 *   	 "abstract":"",   //摘要
 *   	 "author":"2",    //作者
 *   	 "classid":"7",
 *   	 "classname":"图片",
 *   	 "create_time":"1517542845",
 *   	 "del":"0",
 *   	 "id":"3",
 *   	 "imgurl":"http://121.41.116.104:8003/",
 *   	 "name":"1",
 *   	 "remarks":"",
 *       "episodesTotal": "0",   // 漫画集数
 *   	 "uid":"4",
 *   	 "update_time":"1517542845"
 *    },"status":200}
 *
 *
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 *
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://121.41.116.104:8003/api/cartoon/info
 */
func cartoonInfo(c *mygin.IContext) {
	cartoonId := c.Request.FormValue("id")

	data := ddy.CartoonInfo("0", cartoonId)
	if data[1] != nil {
		data[1]["episodesTotal"] = ddy.CartoonEpisodesNum(cartoonId)
	}

	R.Success(c.Writer, data[1])
}


/**
 * @api {get} /api/cartoon/episodesInfo 获取漫画集数详情
 * @apiName 获取漫画集数详情
 * @apiGroup 漫画
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {int} id 漫画id
 * @apiParam {int} episodes 集数
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"data":{
 *      "cartoon_id":"4",
 *		"createtime":"1519458326",
 *		"id":"7",
 *		// imglist 用逗号分隔的图片地址
 *		"imglist":"http://xt.yundapian.com/upload/cartoon/img/2018/February/24/a68261b8fa708f5dc4f1c795a2e1c12e.png,http://xt.yundapian.com/upload/cartoon/img/2018/February/24/0cca5b1408eef89fd7881536849eaeca.png,http://xt.yundapian.com/upload/cartoon/img/2018/February/24/bae87b2b593d0efa2f69d49d51b8bc76.png,http://xt.yundapian.com/upload/cartoon/img/2018/February/24/e7d98a68dbc8c16ed908d266fdb5e7eb.png",
 *		"sort":"13",
 *		"title":"什么什么的东西2",
 *		"uid":"4"
 *    },"status":200}
 *
 *
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 *
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://121.41.116.104:8003/api/cartoon/episodesInfo
 */
func cartoonEpisodesInfo(c *mygin.IContext) {
	cartoonId := c.Request.FormValue("id")
	episodes := c.Request.FormValue("episodes")

	data := ddy.CartoonEpisodesForNum(cartoonId, episodes)

	R.Success(c.Writer, data)
}