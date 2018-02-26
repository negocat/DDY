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
	// "strconv"
)

/**
 * @api {get} /api/cartoon/class 获取漫画分类列表
 * @apiName 获取漫画分类列表
 * @apiGroup 漫画
 *
 * @apiVersion 1.0.0
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功
 * @apiSuccessExample Success-Response:
 * {"data":[
 *     {
 *    	 "create_time":"1503824383",
 *    	 "id":"7",
 *    	 "imgurl":"http://121.41.116.104:8003/",   // 分类icon
 *    	 "level":"1",
 *    	 "list":"7,
 *    	 ",
 *    	 "name":"图片",            // 分类名称
 *    	 "remarks":"",
 *    	 "topid":"7",
 *    	 "uid":"4",
 *    	 "upperid":"0"
 *     }],"status":200}
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
 * @apiSampleRequest http://121.41.116.104:8003/api/cartoon/class
 */
func classList(c *mygin.IContext) {

	data := ddy.ClassForLevel(1, 0)

	R.Success(c.Writer, mysqldb.List(data))
}