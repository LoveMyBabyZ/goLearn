package router

import (
	"avatarzlreportcenter/config"
	"avatarzlreportcenter/controllers/report"
	"avatarzlreportcenter/middlewares/logger"
	sessionMiddleware "avatarzlreportcenter/middlewares/session"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(logger.Logger())
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "welcome to wx")

	})

	router.Use(sessionMiddleware.RegisterSession(config.CookieConfig.Name))

	courseGroup := router.Group("/api/hospital")
	{
		courseGroup.GET("/report-list", report.GetHospitalReportList)
		courseGroup.GET("/online-date", report.GetHospitalOnlineDate)
	}
	//router.GET("/import/import-data",imports.ImportProduct)
	return router
}
