/*
@Time : 2021/12/9 20:47
@Author : 康
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
	Hobby  string
}

func main() {
	//GORM(不需要写sql语句，但是性能没原生sql语句好)
	//连接数据库

	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gostudy?charset=utf8")
	//一个坑，不设置这个参数，gorm会把表名转义后加个s，导致找不到数据库的表
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	//创建表，自动迁移（结构体和数据对应）妙啊！！！！
	db.AutoMigrate(&UserInfo{})
	u := UserInfo{"kzr", "boy", 23, "game"}
	db.Create(&u)
	// 查询
	var u2 = new(UserInfo)
	db.First(u2) //查询并保存在u2中
	fmt.Printf("%#v\n", u2)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u2).Update("hobby", "双色球")
	// 删除
	db.Delete(&u2)

}
