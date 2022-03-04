/*
@Time : 2021/11/24 12:54
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
	r := gin.Default()
	r.LoadHTMLFiles("templates/login.html",
		"templates/index.tmpl")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//method1
		//name := c.PostForm("username")
		//passowrd := c.PostForm("password")
		//method2
		//name := c.DefaultPostForm("username", "xxx")
		//password := c.DefaultPostForm("password", "xxxxxx")
		//method3
		name, ok := c.GetPostForm("username")
		if !ok {
			name = "xxx"
		}
		//password, ok := c.GetPostForm("password")
		//if !ok {
		//	password="xxxxx"
		//}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": name,
		})
	})
	r.Run(":9090")
}
