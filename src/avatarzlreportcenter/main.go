package main

import (
	"avatarzlreportcenter/config"
	"avatarzlreportcenter/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func main() {
	if config.ServerConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := router.InitRouter()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
