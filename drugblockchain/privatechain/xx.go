/*
@Time : 2021/12/21 18:32
@Author : 康
@File : xx
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	CT     string
	Token  string
	Finger string
}

type drug struct {
	id  int
	msg Message
}

func GetDrugInfoBypPWD() {
}

func GetDrugInfoByID() {

}

func GetDrugInfoByM() {

}
func main() {
	engine := gin.Default()
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "ok",
		})
	})

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	//路由组
	njupt := engine.Group("/block")

	njupt.GET("/drug", func(c *gin.Context) {
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
	GetDrugInfoBypPWD()
	GetDrugInfoByID()
	GetDrugInfoByM()
	engine.Run()
}
