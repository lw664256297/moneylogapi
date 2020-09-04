package main

import (
	"errors"
	"log"
	"moneylogapi/router"
	"net/http"
	"time"

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

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	// 监听端口
	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}

// pingServer 检查系统是否正常
func pingServer() error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
