//package course
//
//import (
//	"fmt"
//	"regexp"
//	"strconv"
//	"wx_mini/modules/gredis"
//	"wx_mini/modules/log"
//
//	"wx_mini/config"
//	"wx_mini/modules/httplib"
//	"wx_mini/modules/util"
//)
//
//const cacheCourseList string = "course_list"
//const cacheDetailName string = "course_detail"
//const cacheVisitNum string = "course_visit_num"
//const cacheUserVisitNum string = "user_visit_num"
//
//func GetCourseDetail(id int) (jsonData map[interface{}]interface{}, err error) {
//	var resp string
//	var queryParam map[string]interface{}
//	var returnJsonData map[string]interface{}
//	queryParam["id_course"] = id
//	detailRedisKey := config.CookieConfig.Name + ":"+cacheDetailName
//
//	if gredis.HExists(detailRedisKey,string(id)){
//		bytes, errors := gredis.HGet(detailRedisKey,string(id))
//		if errors != nil {
//			return
//		}
//		resp = string(bytes[:])
//		jsonData, err = jsonToMap(resp)
//	}else{
//		requestInfoUrl := config.HttpRequestConfig.RpCourseList +"course/view"//详情地址
//		requestChapterUrl := config.HttpRequestConfig.RpCourseList +"lesson/index"//章节列表
//		requestRawUrl := config.HttpRequestConfig.RpCourseDetail +"course/raw-info"//课程简介
//		resp,err = httpRequest(requestInfoUrl,queryParam)
//		returnJsonData,err = jsonToMap(resp)
//
//		returnJsonData = (returnJsonData["data"]).(map[string]interface{})
//		sa := make(map[string]string,3)
//		sa["course_info"] = make(map[string]string)
//		jsonData["course_info"] = returnJsonData["list"]
//
//		resp,err = httpRequest(requestChapterUrl,queryParam)
//		returnJsonData,err = jsonToMap(resp)
//		returnJsonData = (returnJsonData["data"]).(map[string]interface{})
//		jsonData["course_chapter"] = returnJsonData["list"]
//
//		resp,err = httpRequest(requestRawUrl,queryParam)
//		returnJsonData,err = jsonToMap(resp)
//		returnJsonData = (returnJsonData["data"]).(map[string]interface{})
//		sa["course_info"]["course_info"] = returnJsonData["list"]
//
//
//
//
//
//	}
//
//
//	return
//}
//
////获取课程列表
//func GetCourseList(action string) (jsonData map[string]interface{}, err error) {
//	var (
//		resp string
//		queryParam map[string]interface{}
//	)
//	requestUrl := config.HttpRequestConfig.RpCourseList + action
//	if gredis.Exists(config.CookieConfig.Name + ":" + cacheCourseList) { //判断是否存在
//		bytes, errors := gredis.Get(config.CookieConfig.Name + ":" + cacheCourseList)
//		log.ErrorLog.Println(errors)
//		if errors != nil {
//			return
//		}
//		resp = string(bytes[:])
//		if resp == "" {
//			resp, err = httpRequest(requestUrl,queryParam)
//		}
//		log.InfoLog.Println(resp)
//	} else {
//		resp, err = httpRequest(requestUrl,queryParam)
//		if err != nil {
//			m := make(map[string]interface{})
//			return m, err
//		}
//		_, err = gredis.Set(config.CookieConfig.Name+":"+cacheCourseList, resp, 3600)
//
//	}
//	log.InfoLog.Println(resp)
//	jsonData, err = util.JsonToMap(resp)
//	log.InfoLog.Println(jsonData)
//	log.ErrorLog.Println(err)
//
//	return
//}
//
////请求
//func httpRequest(requestUrl string, queryParams map[string]interface{}) (resp string, err error) {
//	req, err := httplib.Get(requestUrl)
//	req, err = req.JsonBody(queryParams)
//
//	if err != nil {
//		return
//	}
//
//	resp, err = req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36").String()
//
//	log.InfoLog.Println("[url]=>" + requestUrl)
//	log.InfoLog.Println(resp)
//	if err != nil {
//		return
//	}
//
//	return
//}
//
////json to map
//func jsonToMap(jsonStr string) (jsonData map[string]interface{}, err error) {
//	reg := regexp.MustCompile(`/\*.*\*/`)
//	configStr := reg.ReplaceAllString(jsonStr, "")
//	log.InfoLog.Println(configStr)
//	configStr, err = strconv.Unquote(configStr)
//	bytes := []byte(configStr)
//	if jsonData, err = util.JsonToMap(bytes); err != nil {
//		log.ErrorLog.Println(err)
//	}
//	fmt.Println(jsonData)
//	return
//}
