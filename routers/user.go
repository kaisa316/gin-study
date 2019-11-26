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
	//验证
	if result := user.validate(c); result != nil {
		c.JSON(http.StatusBadRequest, result)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "hello",
	})
}

//验证user
func (u User) validate(c *gin.Context) (errMsg map[string]string) {
	//用户名需要是一个邮箱
	reg, _ := regexp.Compile(`[0-9a-zA-Z]{3,}@\w+\.com$`)
	if reg.MatchString(u.Username) == false {
		errMsg = map[string]string{
			"err_msg": "非法的邮箱地址",
		}
	}
	return
}
