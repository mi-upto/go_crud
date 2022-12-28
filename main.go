package main

import (
	"go_crud/controller"
	"go_crud/middleware"
	"go_crud/model"
	"go_crud/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := service.SqlConnect()
	db.AutoMigrate(&model.User{}) // 構造を DBに反映

	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	userEngine := engine.Group("/user")
	{
		v1 := userEngine.Group("/v1")
		{
			v1.POST("/add", controller.UserAdd)
			v1.GET("/list", controller.UserList)
			v1.PUT("/update/:id", controller.UserUpdate)
			v1.DELETE("/delete/:id", controller.UserDelete)
		}
	}

	engine.Run()

	defer db.Close()
}
