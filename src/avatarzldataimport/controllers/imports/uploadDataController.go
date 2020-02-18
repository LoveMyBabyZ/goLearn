package imports

import (
	"avatarzldataimport/modules/e"
	"avatarzldataimport/modules/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func UploadData(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.InfoLog.Println(file.Filename, file.Size)
	fileName := "文件" + strconv.FormatInt(time.Now().Unix(),10) + ".xlsx"
	err := c.SaveUploadedFile(file, "/tmp/"+fileName)
	if err != nil {
		log.ErrorLog.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": "",
	})
}
