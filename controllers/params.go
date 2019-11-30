package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

//GET :普通形式获取query中的请求参数
func NormalQueryParam(c *gin.Context) {
	name := c.Query("name")
	addr := c.DefaultQuery("address", "海淀区")
	c.String(http.StatusOK, name)
	c.String(http.StatusOK, addr)
}

//GET:bind 结构体方式获取query中的请求参数
func BindQueryParam(c *gin.Context) {
	var person Person
	c.BindQuery(&person) //这个方法会终止请求
	//c.ShouldBindQuery(&person) ,这方法不会终止请求
	c.String(http.StatusOK, person.Name)
	c.String(http.StatusOK, person.Address)
}

//POST=form+query中的请求参数
//bind结构体方式能够获取form+query中的参数
func BindPostformParam(c *gin.Context) {
	var person Person
	c.Bind(&person)
	// c.ShouldBind(&person)//发生错误时，不会终止请求
	c.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Address,
	})
}

//POST: 普通形式获取post请求参数
//这种形式只能后去form中的参数，不能获取query中的参数
func NormalPostformParam(c *gin.Context) {
	name := c.PostForm("name")
	addr := c.PostForm("address")
	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"address": addr,
	})
}

func ParamsInPath(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func ParamsInallPath(c *gin.Context) {
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": action,
	})
}
