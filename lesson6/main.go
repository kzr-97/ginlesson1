/*
@Time : 2021/12/7 17:43
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	//重定向
	engine.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	engine.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		engine.HandleContext(c)
	})
	//跳转
	engine.GET("/test3", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "test3",
		})

	})
	engine.Run(":9090")
}
