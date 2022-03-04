/*
@Time : 2021/12/5 13:46
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
	engine.GET("/blog/:year/:month", func(c *gin.Context) {
		//获取路径参数
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	engine.Run(":9090")
}
