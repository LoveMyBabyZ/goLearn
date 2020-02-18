package router

import (
	"avatarzldataimport/config"
	"avatarzldataimport/controllers/imports"
	"avatarzldataimport/middlewares/logger"
	sessionMiddleware "avatarzldataimport/middlewares/session"
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

	importGroup := router.Group("/api/web")
	{
		importGroup.POST("/upload/data",imports.UploadData)
		importGroup.POST("/import/import-data",imports.ImportData)
	}
	return router
}
