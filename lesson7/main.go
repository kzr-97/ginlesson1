/*
@Time : 2021/12/7 18:31
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
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "ok",
		})
	})
	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404代码，下面的代码为没有匹配到路由的请求都返回views/404.html页面。
	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	//路由组
	njupt := engine.Group("/njupt")

	njupt.GET("/mei", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "mei",
		})
	})
	njupt.GET("/lan", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "lan",
		})
	})
	njupt.GET("/zhu", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "zhu",
		})
	})
	njupt.GET("/ju", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ju",
		})
	})

	engine.Run()
}
