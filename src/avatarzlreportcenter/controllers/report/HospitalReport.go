package report

import (
	"avatarzlreportcenter/models/hospital"
	"avatarzlreportcenter/modules/e"
	"avatarzlreportcenter/modules/log"
	"avatarzlreportcenter/services"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHospitalReportList(c *gin.Context) {
	where := make(map[string]interface{})
	where["org_code"] = c.Query("org_code")
	where["report_date"] = c.Query("date")
	data, err:= services.GetHospitalReport(where)
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
		"data":    data,
	})

}

func GetHospitalOnlineDate(c *gin.Context) {
	where := make(map[string]interface{})
	where["org_code"] = c.Query("org_code")
	if where["org_code"] == "" {
		c.JSON(e.INVALID_PARAMS, gin.H{
			"code":e.INVALID_PARAMS,
			"msg": e.GetMsg(e.INVALID_PARAMS),

		})
		return
	}
	data, err := hospital.GetByCodeOrId(where)
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
		"data":    data,
	})
}
