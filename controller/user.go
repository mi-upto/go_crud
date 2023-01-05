package controller

import (
	"go_crud/model"
	"go_crud/service"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService := service.UserService{}
	err = userService.SetUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func UserList(c *gin.Context) {
	userService := service.UserService{}
	userLists := userService.GetUserList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"users":   userLists,
	})
}

func UserUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	user := model.User{}

	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	userService := service.UserService{}
	err = userService.UpdateUser(int(id), &user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func UserDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	userService := service.UserService{}
	err = userService.DeleteUser(int(id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
