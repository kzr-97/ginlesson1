/*
@Time : 2021/9/18 10:58
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Msg struct {
	//大写才能被访问，若json格式为小写，如何统一？用tag
	Name    string `json:"name"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

func main() {
	r := gin.Default()
	//加在html渲染模板，可以加载多个
	r.LoadHTMLFiles("index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		//http请求
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "liwenzhou.com",
		})
	})
	r.GET("/hello", sayHello)
	r.GET("/json", func(c *gin.Context) {
		data := Msg{
			Name:    "kzr",
			Message: "hello",
			Age:     16,
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/web", func(c *gin.Context) {
		//http://localhost:9090/web?query=kzr&age=18
		//？后面是querystring参数，多个key-value用&连接
		name := c.Query("query")           //通过query请求参数从浏览器获取数据到后端
		age := c.DefaultQuery("age", "20") //another method
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")

}
func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello go",
	})

}
