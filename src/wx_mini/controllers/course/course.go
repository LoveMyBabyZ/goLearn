package course

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wx_mini/models/course"
	"wx_mini/modules/log"
)

var jsonData map[string]interface{}

func CourseList(c *gin.Context) {
	jsonData, err := course.GetCourseList("course/list")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "json convert error",
		})
		log.ErrorLog.Println(err)
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "",
		"data":    jsonData["data"],
	})
}
func CourseDetail(c *gin.Context) {
	id := c.Param("id_course")
}
