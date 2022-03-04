/*
@Time : 2021/12/5 14:50
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//lesson4:gin参数绑定（之前所有参数都是一个一个传见lesson1，这里直接用shouldbind方法绑定，自动识别参数利用反射识别对应结构体）
type UserInfo struct {
	//此处用大写，因为后续shouldbind涉及到反射，因此要大写才能获取参数
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	engine := gin.Default()
	engine.GET("login", func(c *gin.Context) {
		var u UserInfo
		// ShouldBind()会根据请求的Content-Type(url参数、form表单、json数据都行)自行选择绑定器
		//表述可能有误，待定
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	engine.POST("form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	engine.Run(":9090")
}
