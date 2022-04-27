package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

//路由入口文件
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	// 后端管理路由接口
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		// 修改密码
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)
		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCategory)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArticleInfo)
		auth.GET("admin/article", v1.GetArticle)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.UpLoad)

		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)

		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	// 前端登录路由模块
	router := r.Group("api/v1")
	{
		// 用户模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类模块
		router.GET("category", v1.GetCategory)
		router.GET("category/:id", v1.GetCategoryInfo)

		// 文章模块
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCategoryArticle)
		router.GET("article/:id", v1.GetArticleInfo)

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		//评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)

	}
	_ = r.Run(utils.HttpPort)
}
