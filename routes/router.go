package routes

import (
	"ginbolg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由入口文件

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("Hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(utils.HttpPort)
}
