/*
@Time : 2021/12/8 16:14
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//中间件（本质是AOP？）
func main() {
	engine := gin.Default()
	engine.Use(TimeCost())
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	//路由组使用中间件
	test2 := engine.Group("/test2")
	test2.Use(TimeCost())
	//method2：test2 := engine.Group("/test2", TimeCost())
	{
		//...
	}
	engine.Run()
}
func TimeCost() gin.HandlerFunc { //用函数包装
	//在这里调用数据库，查询条件之类
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		fmt.Println("timecost:", cost)
	}
}
