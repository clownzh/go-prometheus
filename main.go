package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//初始化gin
	router := gin.Default()
	//pprof 性能分析
	ginpprof.Wrapper(router)

	//普罗米修斯
	router.Use(ginprom.PromMiddleware(nil))
	// register the `/metrics` route.
	router.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	router.Run(":8080")
}
