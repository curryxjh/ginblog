package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//后端登录
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	data, code = model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCSE {
		setToken(c, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data.Username,
			"id":      data.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// 前台登录
func LoginFront(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	var code int
	data, code = model.CheckLoginFront(data.Username, data.Password)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data.Username,
		"id":      data.ID,
		"message": errmsg.GetErrMsg(code),
	})
}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg((errmsg.ERROR)),
			"token":   token,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
	return
}
