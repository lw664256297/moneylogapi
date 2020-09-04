package main

import (
	"log"
	"moneylogapi/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin框架路由
	g := gin.New()

	// 初始化中间件
	middleware := []gin.HandlerFunc{}

	// 初始路由
	router.Load(
		g,
		middleware...,
	)

	// 监听端口
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g.Error()))

}
