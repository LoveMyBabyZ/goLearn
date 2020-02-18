package main

import (
	"Server/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"wx_mini/router"
)

func main() {
	if config.ServerConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := router.InitRouter()
	s := &http.Server{
		Addr:        ":" + strconv.Itoa(config.ServerConfig.Port),
		Handler:     router,
		ReadTimeout: 15 * time.Second,
		MaxHeaderBytes:1<<20,
	}
	s.ListenAndServe()
}
