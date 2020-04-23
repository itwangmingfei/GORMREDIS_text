package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"text/tools"
)

//定义一个结构GORM构造数据库
type Users struct {
	gorm.Model
	Username string `gorm:"size:15;commit:'姓名'"`
	Age  int `gorm:"commit:'年龄'"`
}


func main(){
	r :=gin.Default()
	ph,_ := os.Getwd()
	path := "/config/config.json"
	//初始化配置文件
	tools.ReadConfig(ph + path)
	//初始化DB
	_,err :=tools.InitGorm()
	if err!= nil{
		panic(err)
	}
	//**********************************
	//实例化表迁移创建表
	tools.DB.AutoMigrate(&Users{})
	//********简单使用GORM
	//var user  Users

	//添加----------------------------
	//user.Age=12
	//user.Username ="Giao"
	//tools.DB.Create(&user)
	//修改----------------------------1
	//user.ID = 2
	//user.Username="小明"
	//user.Age = 20
	//tools.DB.Model(&user).Update(&user)
	//--------------------------------2
	//tools.DB.Where("id = ?",2).First(&user)
	//user.Age=18
	//user.Username = "lucy"
	//tools.DB.Model(&user).Update(&user)
	//--------------------------------3
	//tools.DB.Model(&user).Where("id =?",4).Update("age",13)
	//删除-----------------------------------1
	//tools.DB.Where("id = ?",2).First(&user)
	//tools.DB.Model(&user).Delete(&user)
	//---------------------------------------2
	//tools.DB.Where("id=?","5").Delete(&user)

	//查看一条数据
	//var suser Users
	//tools.DB.Where("id = ?",2).First(&suser)
	//******************************
	//查看分页数据
	//var aluser []Users
	//tools.DB.Limit(2).Offset(1).Order("id desc").Find(&aluser)


	//初始化redis
	//-------------------------------------------------------------------------------
	tools.InitRedisStore()
	tools.GOREDIS.Set("wmf","ggglook")


	r.GET("/", func(ctx *gin.Context) {
		//显示json文件


		ctx.JSON(200,gin.H{"hellow":tools.Cfg,"key":"str"})
		//初始化DB
	})

	//使用配置接口
	port := tools.Cfg.AppPort
	if port!=""{
		apphost :=tools.Cfg.AppHost
		panic(r.Run(apphost+":"+port))
	}
	panic(r.Run())
}
