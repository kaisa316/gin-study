package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaisa316/gin-study/models"
	"regexp"

	"net/http"
)

type User struct {
	// Username string `form:"username" binding:"required"`
	// Password string `form:"password" binding:"required"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func CreateTable(c *gin.Context) {
	models.CreateTable()
	c.JSON(http.StatusOK, gin.H{
		"msg": "创建表成功",
	})
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//验证
	if errMsg := user.validate(c); errMsg != nil {
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	userModel := models.User{
		Username: user.Username,
		Password: user.Password,
	}
	models.AddRecord(&userModel)

	//返回
	c.JSON(http.StatusOK, map[string]string{
		"errno": "0",
		"msg":   "注册成功",
	})
}

//查询User信息
func QueryUser(c *gin.Context) {
	result := models.InfoByUsername()
	c.JSON(http.StatusOK, result)
}

//验证user
func (u User) validate(c *gin.Context) (err map[string]string) {
	err = make(map[string]string)
	//用户名需要是一个邮箱
	reg, _ := regexp.Compile(`[0-9a-zA-Z]{3,}@\w+\.com$`)
	if reg.MatchString(u.Username) == false {
		err["errno"] = "1001"
		err["msg"] = "非法的邮箱地址"
		return
	}
	reg, _ = regexp.Compile(`.{6,}`)
	if reg.MatchString(u.Password) == false {
		err["errno"] = "1002"
		err["msg"] = "非法的密码"
		return
	}

	return nil
}
