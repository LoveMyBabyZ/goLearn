package router

import (
	"github.com/gin-gonic/gin"
	"wx_mini/config"
	"wx_mini/controllers/course"
	"wx_mini/middlewares/logger"
	sessionMiddleware "wx_mini/middlewares/session"
	"wx_mini/modules/gredis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(logger.Logger())
	router.Use(gin.Recovery())
	gredis.Setup()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "welcome to wx")

	})

	router.Use(sessionMiddleware.RegisterSession(config.CookieConfig.Name))

	courseGroup := router.Group("/api/course")
	{
		courseGroup.GET("/list", course.CourseList)
	}

	return router
}
