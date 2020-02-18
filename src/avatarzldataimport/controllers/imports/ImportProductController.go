package imports

import (
	"avatarzldataimport/modules/e"
	"avatarzldataimport/modules/log"
	"avatarzldataimport/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ImportData(c *gin.Context) {
	//params := make(map[string]interface{})
	//params["path"] = "/Users/zhaodeyong/Desktop/商品新增.xlsx"
	//params["proudct_type"] = 1010
	//fmt.Print(params)
	ImportHandler := services.Params{}
	ImportHandler.Path = c.PostForm("url_path")
	if c.PostForm("url_path") == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_PARAM,
			"msg":  "上传的路径不能为空",
			"data": "",
		})
	}
	ImportHandler.Type, _ = strconv.Atoi(c.PostForm("type"))
	if c.PostForm("type") == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_PARAM,
			"msg":  "类型不能为空",
			"data": "",
		})
	}
	fmt.Println(ImportHandler.Type)
	fmt.Println(ImportHandler.Path)
	msg := ImportHandler.ImportProduct()

	if msg != "" {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_DATA,
			"msg":  msg,
			"data": "",
		})
		log.ErrorLog.Println(msg)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": "",
	})
}
