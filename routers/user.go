package routers

import (
	"github.com/gin-gonic/gin"
	"regexp"

	"net/http"
)

type User struct {
	// Username string `form:"username" binding:"required"`
	// Password string `form:"password" binding:"required"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	reg, _ := regexp.Compile(`[0-9a-zA-Z]{3,}@\w+=\.com$`)
	//用户名需要时一个邮箱
	if reg.MatchString(user.Username) == false {

		c.JSON(http.StatusBadRequest, gin.H{
			"err_msg": "非法的邮箱地址",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "hello",
	})
}
